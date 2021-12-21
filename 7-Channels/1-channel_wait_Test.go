package main

import (
	"fmt"
	"time"
)



func main() {


	ch := make(chan string)

	//ch <- "hello from main"  //This won't work since it blocks
	go sendMe(ch)
	//for i:=1; i<2; i++{  // This for loop just reads the channel as messages come in.
	fmt.Println(<-ch)
	//}


}
func sendMe(ch chan<- string) {

	time.Sleep(time.Second*2)
	ch <- "SendMe is done"


}
