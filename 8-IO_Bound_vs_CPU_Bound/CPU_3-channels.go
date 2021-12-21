// This is concurrency / goroutines implemented with CHANNELS
package main

import (
	"fmt"
	"runtime"

	"time"
)



func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(16)   //  Extra processors help up to number of goroutines w CPU-bound tasks
	c := make(chan string)
	startTime := time.Now()
	go counta(c)
	go countb(c)
	go countc(c)
	go countd(c)
	go counte(c)
	go countf(c)
	go countg(c)
	go counth(c)



	for i := 0; i < 8; i++ {
		fmt.Println(<-c)
		}

	elapsed := time.Since(startTime)
	fmt.Printf("Processes took %s", elapsed)
}
func counta(c chan string) {
	fmt.Println("AAAA is starting  ")
	for I := 1; I < 10_000_000_000; I ++ {
	}


	c <- "AAAA is done"

}
func countb(c chan string) {
	fmt.Println("BBBB is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}


	c <- "BBBB is done"

}
func countc(c chan string) {
	fmt.Println("CCCC is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	c <- "CCCC is done"


}
func countd(c chan string) {
	fmt.Println("DDDD is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}


	c <- "DDDD is done"

}
func counte(c chan string) {
	fmt.Println("EEEE is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	c <- "EEEE is done"


}
func countf(c chan string) {
	fmt.Println("FFFF is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}


	c <- "FFFF is done"

}
func countg(c chan string) {
	fmt.Println("GGGG is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	c <- "GGGG is done"


}
func counth(c chan string) {
	fmt.Println("HHHH is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	c <- "HHHH is done"


}