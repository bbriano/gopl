// Argv prints the argument list.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, s := range os.Args {
		fmt.Printf("%d\t%s\n", i, s)
	}
}
