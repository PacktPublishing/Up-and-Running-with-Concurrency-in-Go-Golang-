package main

import (
	"fmt"
	"time"
)

func main() {

	var value = 97
	var result = 0
	goChan := make(chan int)
	mainChan := make(chan string)
	calculateSquare := func() { // This is a nested function, assigned to a variable
		time.Sleep(time.Second * 3)
		result = value * value
		goChan <- result
	}
	reportResult := func() { // This is a nested function, assigned to a variable
		fmt.Println(value, "squared is", <-goChan)
		// blocks until it can read something from goChan - printed
		mainChan <- "You can quit now.  I'm done." // This is just for clarity.
	}

	go calculateSquare()
	go reportResult()
	<-mainChan // blocks until it can read something from mainChan - discarded

}
