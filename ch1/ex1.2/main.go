// ex1.2 prints commandline indexes and arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
	//
	//OR:
	//
	for i, arg := range os.Args {
		fmt.Printf("%d: %s\n", i, arg)
	}
}
