//	This file is a complete, operating, SYNCHRONOUS (non-concurrent) program.
//	Your mission is to cut execution time in half by making adding concurrency.
//	Before you start, close all other open files in your IDE.
//	1) Run this code now and verify that it works with no errors.  Note the elapsed time.
//	2) Notice the commands and keywords below related to concurrency
// 	wg.Done(), wg.Add(2), import "sync", wg.Wait(), go, var wg = sync.WaitGroup{}
//	3) Insert the commands above at the appropriate places below to concurrently run doSomething and doSomethingElse.  Some commands or keywords may be needed more than once.
//	4) Run the code with your changes and verify correct operation.  Note the elapsed time.
//	5) Decrease the sleep time on doSomethingElse to 1000 ms. How will this effect output and elapsed time?
//	6) Run the code with your changes and verify correct operation.  Note the elapsed time.
//	7) Were your predictions correct?
//	8) Make other changes and experiment as you like. Take your time and be creative!
//	9) Open SOLUTION-UsingWaitgroups.go if you get stuck
//	10) Share any interesting discoveries.

package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	doSomething()
	doSomethingElse()
	fmt.Println("I guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func doSomething() {
	time.Sleep(time.Millisecond * 1500)
	fmt.Println("\nI've done something")
}

func doSomethingElse() {
	time.Sleep(time.Millisecond * 1500)
	fmt.Println("I've done something else")
}
