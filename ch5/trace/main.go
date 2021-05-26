// The trace program uses defer to add entry/exit diagnostics to a function.
package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	//
	// Don't forget the FINAL PARENTHESES,
	// or the on-entry action will happen on exit
	// and the on-exit action won't happen at all.
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	//
	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func main() {
	bigSlowOperation()
}

/*
!+output
$ go build gopl.io/ch5/trace
$ ./trace
2015/11/18 09:53:26 enter bigSlowOperation
2015/11/18 09:53:36 exit bigSlowOperation (10.000589217s)
!-output
*/
