package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 5000, "socks": 500}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

type cents int

func (c cents) String() string {
	return fmt.Sprintf("$%d.%02d", c/100, c%100)
}

type database map[string]cents

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
