// Xkcd search for comic with regular expression and prints link.
package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

//go:embed xkcd.json
var comicsJson []byte

type Comic struct {
	Num        int
	Title      string
	Transcript string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: xkcd regexp\n")
		os.Exit(1)
	}

	var comics []Comic

	err := json.Unmarshal(comicsJson, &comics)
	if err != nil {
		fmt.Fprintf(os.Stderr, "xkcd: %v\n", err)
		os.Exit(1)
	}

	re, err := regexp.CompilePOSIX(strings.Join(os.Args[1:], " "))
	if err != nil {
		fmt.Fprintf(os.Stderr, "xkcd: %v\n", err)
		os.Exit(1)
	}

	for _, c := range comics {
		text := strings.ToLower(c.Title + " " + c.Transcript)
		if re.MatchString(text) {
			fmt.Printf("https://xkcd.com/%d\n", c.Num)
		}
	}
}
