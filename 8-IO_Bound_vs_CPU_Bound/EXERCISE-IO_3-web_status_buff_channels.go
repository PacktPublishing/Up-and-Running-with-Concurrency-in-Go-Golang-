//	UNDERSTANDING the LEN AND RANGE COMMANDS WITH CONCURRENCY
//	This file checks multiple websites concurrently.
//	Your mission is to understand two ways to ensure that all links are checked.
//	Before you start, close all other open files in your IDE.
//	When operating correctly, this program should check 7 links and report their status.
// 	Note that the results will not be in order.
//	1) Run this code now and verify correct operation.
//	2) Examine the code starting with "for len(c) < len(links)".  What does this do?  You may need to refer to documentation for len.
//	3) Comment out the for loop starting with "for len(c) < len(links)".  Be sure to include the { }
//	4) UN Comment the code section that starts with "for range (links)".  Run again and verify proper operation.
//	5) Note that this other approach also includes, "channel message".
// 	6) What does "for range (links)" do? Which is the better approach, using len(c) as above, or range(links)?  Why?
//	7) Make other changes and experiment as you like. Take your time and be creative!
//	8) Open SOLUTION-IO_3-web_status_buff_channels.go for full solution.
//	9) Share any interesting discoveries.

package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)



func main() {
	runtime.GOMAXPROCS(1)

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://cnn.com",
		"http://udemy.com",
	}
	c := make(chan string, len(links))

	start := time.Now()

	for _, link := range links {
		go checkLink(link, c)
	}
	for len(c) < len(links) { // Program ends when everything is in BUFFERED channel.
	}
	//for range links{  // This also works beautifully, using blocking code
	//	fmt.Println("channel message:",<-c)
	//}

	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
