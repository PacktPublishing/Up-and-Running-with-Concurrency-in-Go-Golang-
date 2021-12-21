// This is concurrency / goroutines implemented with a WAITGROUP
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(16)     //  Extra processors help up to number of goroutines w CPU-bound tasks
	wg.Add(8)  // create 8 contractor badges for our workers
	startTime := time.Now()
	go counta() // pass out a badge to each worker
	go countb()
	go countc()
	go countd()
	go counte()
	go countf()
	go countg()
	go counth()





	wg.Wait()  // Do not end the program until all badges have been returned, ie all go routines have reported that they are done.
	elapsed := time.Since(startTime)
	fmt.Printf("Processes took %s", elapsed)
}
func counta() {
	fmt.Println("AAAA is starting  ")
	for I := 1; I < 10_000_000_000; I ++ {
	}

	fmt.Println("AAAA is done  ")
	wg.Done()  // Turn in my badge - I'm done

}
func countb() {
	fmt.Println("BBBB is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("BBBB is done")
	wg.Done()

}
func countc() {
	fmt.Println("CCCC is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("CCCC is done    ")
	wg.Done()

}
func countd() {
	fmt.Println("DDDD is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("DDDD is done     ")
	wg.Done()

}
func counte() {
	fmt.Println("EEEE is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("EEEE is done   ")
	wg.Done()

}
func countf() {
	fmt.Println("FFFF is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("FFFF is done     ")
	wg.Done()

}
func countg() {
	fmt.Println("GGGG is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("GGGG is done     ")
	wg.Done()

}
func counth() {
	fmt.Println("HHHH is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("HHHH is done     ")
	wg.Done()

}