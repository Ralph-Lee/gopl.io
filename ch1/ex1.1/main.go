// ex1.1 prints command line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args)
	// OR:
	fmt.Println(strings.Join(os.Args[0:], " "))
}
