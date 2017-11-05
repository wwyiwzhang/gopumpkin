package main

import (
	"fmt"
	"time"
	"sync"
)

// pizza tracker
var wg sync.WaitGroup

func main() {

	orders := []int{1, 2, 3, 4, 5, 6}
	orders_chan := make(chan int)
	wg.Add(2)
	go make_pizza(orders, orders_chan)
	go deliver_pizza(orders_chan)
	//time.Sleep(20 * time.Second) // had to add this because if not, main exit before goroutine finishes executing
	wg.Wait()
	fmt.Println("All delivery completed")
}

func make_pizza(orders []int, c chan int) {
	defer wg.Done()
	for _, val := range orders {
		time.Sleep(100 * time.Millisecond) // the time to make a pizza
		fmt.Println("Completed order number", val)
		c <- val
	}
	close(c)
}

func deliver_pizza(c chan int) {
	defer wg.Done()
	for val := range c {
		time.Sleep(100 * time.Millisecond) // the time to deliver a pizza
		fmt.Println("Order number", val, "has been delivered")
	}
}
