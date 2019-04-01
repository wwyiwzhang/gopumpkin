/* semaphore differentiates from mutex at
it set how many goroutines can access the
code at the same time while mutex limit that
there can be only one goroutine which can
access the code at one time
*/

package main

import (
	"flag"
)

func process(jobs int, size int) {
	init := 0
	sem := make(chan int, size)

	for init < jobs {
		sem <- 1
		go func(init int) {
			// fmt.Printf("processing job %d\n", init)
			// time.Sleep(100 * time.Millisecond)
			<-sem
		}(init)
		init++
	}
}

func main() {
	jobs := flag.Int("j", 100, "number of jobs submitted")
	sn := flag.Int("s", 10, "maximum concurrent access")
	flag.Parse()
	process(*jobs, *sn)
}
