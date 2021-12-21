//	USING CHANNELS INSTEAD OF SHARING VARIABLES (widgetInventory in this case)
// SOLUTION FILE
package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
	//mutex = sync.Mutex{}
	widgetInventory int32 = 1000 //Package-level variable will avoid all the pointers
	c = make(chan int32, 6000) // Creates a buffered channel to hold all our transactions
)

func main() {
	fmt.Println("Starting inventory count = ", widgetInventory)
	wg.Add(2)
	go makeSales(c)
	go newPurchases(c)
	// Put -100 transaction into the channel
	wg.Wait() // Don't proceed until all transactions are loaded into channel
	for len(c) > 0 { // Keep reading until all transactions are processed and channel is empty
		widgetInventory += <-c // Read each transaction and update widgetInventory
		//fmt.Println(widgetInventory) // Print if you want to see each transaction
	}
	fmt.Println("Ending inventory count = ", widgetInventory)
}

func makeSales(c chan int32) { // 300000 widgets sold
	for i := 0; i < 3000; i++ {
		//mutex.Lock()
		//widgetInventory -= 100
		////fmt.Println(widgetInventory)
		//mutex.Unlock()
		c <- -100 // Put -100 sales transaction into the channel
	}

	wg.Done()
}

func newPurchases(c chan int32) { // 300000 widgets purchased
	for i := 0; i < 3000; i++ {
		//mutex.Lock()
		//widgetInventory+= 100
		////fmt.Println(widgetInventory)
		//mutex.Unlock()
		c <- 100  // Put +100 purchase transaction into the channel
	}
	wg.Done()
}
