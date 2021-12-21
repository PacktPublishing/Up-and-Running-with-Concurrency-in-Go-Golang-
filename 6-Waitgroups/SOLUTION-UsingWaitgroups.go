package main

import (
	"fmt"
	"sync"
	"time"
)
var wg = sync.WaitGroup{}  // This creates our waitgroup called, "wg"

func main(){

	wg.Add(2)  // This sets our wait group to 2, meaning we will wait for 2 goroutines to complete
	start := time.Now()
	go doSomething()  // Adding the keyword "go" creates a goroutine
	go doSomethingElse()  // Adding the keyword "go" creates a goroutine
	wg.Wait()  // This blocks (waits) until wg=0 meaning all goroutines are done.
	fmt.Println("I guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func doSomething(){
	time.Sleep(time.Millisecond*1500)
	fmt.Println("\nI've done something")
	wg.Done()  // This decrements wg by one, indicating that doSomething is done.
}

func doSomethingElse(){
	time.Sleep(time.Millisecond*1500)
	fmt.Println("I've done something else")
	wg.Done()  // This decrements wg by one, indicating that doSomethingElse is done.
}
}