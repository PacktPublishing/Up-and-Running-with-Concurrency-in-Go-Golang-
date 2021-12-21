//
//  SEE SOLUTION NOTES BELOW
//

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
	//for len(c) < len(links) {
	//// Stay here until the number of items in channel c = the number of links to be checked
	//// Channel c is a buffered channel that is loaded with each link as it is checked.
	//}
	for range links{  // This also works beautifully, using blocking code.  It stays here until every message in the channel is read.
		fmt.Println("channel message:",<-c)
	}

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
