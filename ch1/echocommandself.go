// Echo prints command itself info.
package main

import (
	"fmt"
	"os"
)

func main() {
	//
	fmt.Println()
	fmt.Print("len(os.Args) -- ")
	fmt.Println(len(os.Args))
	//
	fmt.Println()
	fmt.Print("os.Args[0] -- ")
	fmt.Println(os.Args[0])
	//
	fmt.Println()
	fmt.Print("os.Args -- ")
	fmt.Println(os.Args)
	//
	fmt.Println()
	fmt.Print("os.Args[:len(os.Args)] -- ")
	fmt.Println(os.Args[:len(os.Args)])
	//
	fmt.Println()
	fmt.Print("os.Args[1:len(os.Args)] -- ")
	fmt.Println(os.Args[1:len(os.Args)])
	//
	fmt.Println()
	fmt.Print("os.Args[1:] -- ")
	fmt.Println(os.Args[1:])
}
