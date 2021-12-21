package main

import (
	"fmt"
	"sync"
	)

var (
	wg sync.WaitGroup
	mutex = sync.Mutex{}
	widgetInventory int32= 1000  //Package-level variable will avoid all the pointers
)


func main() {
	fmt.Println("Starting inventory count = ", widgetInventory)
	wg.Add(2)
	go makeSales()
	go newPurchases()
	wg.Wait()
	fmt.Println("Ending inventory count = ", widgetInventory)
}

func makeSales() {  // 1000000 widgets sold
	for i := 0; i < 3000; i++ {
		mutex.Lock()
		widgetInventory -= 100
		fmt.Println(widgetInventory)
		mutex.Unlock()
	}

	wg.Done()
}

func newPurchases() {  // 1000000 widgets purchased
	for i := 0; i < 3000; i++ {
		mutex.Lock()
		widgetInventory+= 100
		fmt.Println(widgetInventory)
		mutex.Unlock()
	}
	wg.Done()
}
