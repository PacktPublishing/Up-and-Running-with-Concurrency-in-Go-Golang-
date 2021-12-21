package main

import (
	"fmt"
	"time"
)

func main(){
	doTheFirstThing()
	doTheSecondThing()
	fmt.Println("No more blocking.  I'm done")
}
func doTheFirstThing() {
	fmt.Println("FirstThing 'blocking' for 2 seconds")
	time.Sleep(time.Second*2)
}

func doTheSecondThing() {
	fmt.Println("SecondThing 'blocking' for 3 seconds")
	time.Sleep(time.Second*3)
}
