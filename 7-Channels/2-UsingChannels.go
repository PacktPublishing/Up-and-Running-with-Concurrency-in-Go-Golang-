package main

import (
	"fmt"
	"time"

)
var ch = make(chan string)

func main(){


	start := time.Now()
	go doSomething()
	go doSomethingElse()


	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("I guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func doSomething(){
	time.Sleep(time.Second*2)
	fmt.Println("\nI've done something")
	ch <- "doSomething finished"
}

func doSomethingElse(){
	time.Sleep(time.Second*2)
	fmt.Println("I've done something else")
	ch <- "doSomethingElse finished"
}