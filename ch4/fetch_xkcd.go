// Fetch Xkcd comic transcript from xkcd.com and prints data in JSON format.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Comic struct {
	Num        int
	Title      string
	Transcript string
}

var comics []Comic

func main() {
	for i := 1; i <= 2582; i++ {
		fmt.Fprintf(os.Stderr, "Fetching comic %d...\n", i)
		c, err := fetchComic(i)
		if err != nil {
			continue
		}
		comics = append(comics, c)
	}

	data, err := json.Marshal(comics)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}

func fetchComic(number int) (Comic, error) {
	var c Comic

	resp, err := http.Get(fmt.Sprintf("https://xkcd.com/%d/info.0.json", number))
	if err != nil {
		return c, err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	io.Copy(&buf, resp.Body)
	err = json.Unmarshal(buf.Bytes(), &c)
	if err != nil {
		return c, err
	}

	return c, nil
}
