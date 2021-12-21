package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10000; i++ {  // This allows us to create MANY goroutines
		wg.Add(1)
		go func(j int) {

			var result = 0
			goChan := make(chan int)
			mainChan := make(chan string)
			calculateSquare := func() {
				time.Sleep(time.Second * 3)  // Deliberate time delay
				result = j * j
				goChan <- result
			}
			reportResult := func() {
				fmt.Println(j, "squared is", <-goChan)
				// blocks until it can read something from goChan - printed
				mainChan <- "You can quit now.  I'm done." // This is just for clarity.
			}

			go calculateSquare()
			go reportResult()
			<-mainChan // blocks until it can read something from mainChan - discarded
			wg.Done()
		}(i)
	}
	wg.Wait()
}
