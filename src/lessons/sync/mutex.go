/*
RWMutex from sync library used for reading/writing shared cache/state
RLock - read lock and doesn't block other reads so in theory we can add
		more RLocks when there is a RLock in place
Lock - write lock and block other write locks and all read locks
*/
package main

import (
	"fmt"
	"sync"
)

type UserCounter struct {
	count map[string]int
	sync.RWMutex
}

func NewUserCounter() *UserCounter {
	return &UserCounter{
		count: make(map[string]int),
	}
}

func (uc *UserCounter) Visit(name string) {
	uc.Lock()
	defer uc.Unlock()
	uc.count[name] += 1
}

func (uc *UserCounter) VisitedCount(name string) int {
	uc.RLock()
	defer uc.RUnlock()
	v := uc.count[name]
	return v
}

func (uc *UserCounter) Count(wg *sync.WaitGroup, userList []string, times int) map[string]int {
	i := 0
	for i < times {
		for _, name := range userList {
			wg.Add(1)
			go func(n string) {
				defer wg.Done()
				uc.Visit(n)
				fmt.Printf("%s visited: %d times\n", n, uc.VisitedCount(n))
			}(name)
		}
		i++
	}
	wg.Wait()
	return uc.count
}

func main() {
	var wg sync.WaitGroup
	newCounter := NewUserCounter()
	fc := newCounter.Count(&wg, []string{
		"user1",
		"user2",
		"user3",
	}, 10)
	fmt.Printf("Final count: %v\n", fc)
}
