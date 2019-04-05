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

type Pool struct {
	workPool []*Worker
}

func newPool() *Pool {
	return &Pool{
		workPool: []*Worker{},
	}
}

func (p *Pool) addWorker(wg *sync.WaitGroup, w *Worker) *Pool {
	wg.Add(1)
	p.workPool = append(p.workPool, w)
	return p
}

func (w *Worker) Do(wg *sync.WaitGroup) {
	defer wg.Done()
loop:
	for {
		if wrk, ok := <-w.wc; !ok {
			break loop
		} else {
			fmt.Printf("Worker #%d is processing work: %v\n", w.id, wrk)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	pool := newPool()
	for i := range []int{1, 2, 3, 4, 5} {
		worker := newWorker(i)
		pool.addWorker(&wg, worker)
		go worker.Do(&wg)
	}

	for j := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		w := Work(j)
		mod := j % 5
		pool.workPool[mod].wc <- w
	}

	go func() {
		for _, w := range pool.workPool {
			fmt.Printf("closing channel for worker: #%d\n", w.id)
			close(w.wc)
		}
	}()

	// Use Wait() function to block the main goroutine
	wg.Wait()
}
