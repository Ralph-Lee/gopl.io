// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//fmt.Println(os.Args[0])
	fmt.Println(os.Args[1:])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
