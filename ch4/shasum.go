// Shasum prints the hash of the standard input.
package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var hashFunction = flag.String("a", "256", "256|384|512")

func main() {
	flag.Parse()
	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)

	switch strings.ToLower(*hashFunction) {
	case "256":
		fmt.Printf("%x\n", sha256.Sum256(buf.Bytes()))
	case "384":
		fmt.Printf("%x\n", sha512.Sum384(buf.Bytes()))
	case "512":
		fmt.Printf("%x\n", sha512.Sum512(buf.Bytes()))
	default:
		fmt.Fprintf(os.Stderr, "invalid hash function\n")
		os.Exit(1)
	}
}
