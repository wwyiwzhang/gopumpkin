/*
This is a simple example of distributing work to multiple workers (goroutines)
Each worker is a struct with an id and an exclusive channel
In the main function, first create five workers (goroutines) and start them then
based on the calculated mod, each Work is sent to the corresponding Worker to process
*/

package main

import (
	"fmt"
	"sync"
)

type Worker struct {
	id int
	wc chan Work
}

type Work int

func newWorker(wid int) *Worker {
	return &Worker{
		id: wid,
		wc: make(chan Work),
	}
}

func (w *Worker) Do(wg *sync.WaitGroup) {
	for {
		wrk, ok := <-w.wc
		if !ok {
			wg.Done()
			break
		}
		fmt.Printf("Worker: #%d is processing work: %v\n", w.id, wrk)
	}
}

func main() {
	var wg sync.WaitGroup
	pool := []*Worker{}

	for i := range []int{1, 2, 3, 4, 5} {
		worker := newWorker(i)
		pool = append(pool, worker)
		wg.Add(1)
		go worker.Do(&wg)
	}

	for j := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		w := Work(j)
		mod := j % 5
		pool[mod].wc <- w
	}

	go func() {
		for _, w := range pool {
			fmt.Printf("closing channel for worker: #%d\n", w.id)
			close(w.wc)
		}
	}()

	// Use Wait() function to block the main goroutine
	wg.Wait()
}
