// Outline prints an outline of an HTML document.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	}
	outline(doc, 0)
}

func outline(n *html.Node, depth int) {
	if n.Type == html.ElementNode {
		for i := 0; i < depth; i++ {
			fmt.Printf("	")
		}
		fmt.Println(n.Data)
	}
	for e := n.FirstChild; e != nil; e = e.NextSibling {
		outline(e, depth+1)
	}
}
