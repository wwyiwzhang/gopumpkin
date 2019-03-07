/*
The context library can be used for exiting all goroutines cleanly
when a request is cancelled and free up resources in a timely manner
This exercise is mimmicking pizza ordering
If receives cancellation request, all the remaining dependent goroutines must return
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func makeDough(ctx context.Context, mc chan int, bc chan int) {
	select {
	case <-ctx.Done():
		fmt.Println("Order canceled!")
	case order := <-mc:
		fmt.Printf("Making dough for order #%d\n", order)
		time.Sleep(100 * time.Millisecond)
		bc <- order
		close(mc)
	}
}

func bakePizza(ctx context.Context, mc chan int, bc chan int) {
	select {
	case <-ctx.Done():
		fmt.Println("Order canceled!")
	case order := <-bc:
		fmt.Printf("Baking pizza for order #%d\n", order)
		time.Sleep(100 * time.Millisecond)
		close(bc)
	}
}

func main() {

	ctx := context.Background()

	cancelCtx, cancelFunc := context.WithCancel(ctx)

	makeChan := make(chan int)
	bakeChan := make(chan int)

	orderNum := 1

	go makeDough(cancelCtx, makeChan, bakeChan)
	go bakePizza(cancelCtx, makeChan, bakeChan)

	makeChan <- orderNum
	cancelFunc()
	time.Sleep(30 * time.Millisecond)
}
