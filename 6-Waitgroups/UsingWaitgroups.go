package main

import (
	"fmt"
	"sync"
	"time"
)
var wg = sync.WaitGroup{}

func main(){

	wg.Add(2)
	start := time.Now()
	go doSomething()
	go doSomethingElse()
	wg.Wait()




	fmt.Println("\n\nI guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func doSomething(){
	time.Sleep(time.Second*2)
	fmt.Println("\nI've done something")
	wg.Done()
}

func doSomethingElse(){
	time.Sleep(time.Second*2)
	fmt.Println("I've done something else")
	wg.Done()
}