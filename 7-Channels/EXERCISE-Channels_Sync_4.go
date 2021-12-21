//	DEBUGGING a File for a Co-worker - THIS MAY BE CHALLENGING!
//	This file attempts to run thousands of simultaneous goroutines, but it has some bugs in it.
//	Your mission is to help your co-worker by identifying and eliminating the bugs.
//	Before you start, close all other open files in your IDE.
//	When operating correctly, this program should square the numbers 1-10,000 and print the result with exit code 0.
// 	Note that the results will not be in order.
//	1) Run this code now and verify the current state.  What error messages did you get?
//	2) Using any debugging techniques you like, correct each bug and rerun the program until it executes correctly.
//	3) If you get stuck, there are hints located at the bottom of this file.  Use only as needed.
//	8) Make other changes and experiment as you like. Take your time and be creative!
//	9) Open SOLUTION-Channels_Sync_4.go for full solution.
//	10) Share any interesting discoveries.


package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10000; i++ {  // This allows us to create MANY goroutines
		wg.Add(3)
		go func(j int) {

			var result = 0
			goChan := make(chan string)
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

}
// HINTS BELOW
//
//
//
//
//
//
//
//
// Hint #1) What type of data will be transmitted in the two channels?
//
// Hint #2) Which statement tells func main() to wait until the goroutines are finished?
//
// Hint #3) How many goroutines do we need to add to our wait group?