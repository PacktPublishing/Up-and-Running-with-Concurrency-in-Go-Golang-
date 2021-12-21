package main

import (
	"fmt"
	"time"
)

var result = 0
var value = 97

func main() {
	goChan := make(chan int)
	mainChan := make(chan string)
	go calculateSquare(value, goChan)
	go reportResult(goChan, mainChan)
	<-mainChan // blocks until it can read something from mainChan - discarded
}
func calculateSquare(value int, goChan chan int) {
	fmt.Println("Calculating for 3 seconds...")
	time.Sleep(time.Second * 3)
	result = value * value
	goChan <- result

}
func reportResult(goChan chan int, mainChan chan string) {
	time.Sleep(time.Second * 1)
	fmt.Println("The result of", value, "squared", "is", <-goChan)
	// blocks until it can read something from goChan - printed
	mainChan <- "You can quit now.  I'm done." // This is just for clarity.
}
