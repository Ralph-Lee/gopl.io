// Dedup prints only one instance of each line; duplicates are removed.
package main

import (
	"bufio"
	"fmt"
	"os"
)

//
//!+
//
func main() {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("~~Notes: The scanner will stop on EOF (End Of File).~~~~~~~~~")
	fmt.Println("~~Typing [Ctrl-D] will send EOF and stop the scanner.~~~~~~~")
	fmt.Println("~~On windows use [Ctrl-Z]~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

//!-
