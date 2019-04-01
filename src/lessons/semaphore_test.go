package main

import (
	"fmt"
	"testing"
)

func BenchmarkSemaphore(b *testing.B) {
	// test if the size of the buffered channel will speed things up and it actually helps to
	// an extend per the following table
	/*
		BenchmarkSemaphore/buffered_channel_size:_10-4         	   10000	   2849151 ns/op
		BenchmarkSemaphore/buffered_channel_size:_100-4        	    5000	   1574254 ns/op
		BenchmarkSemaphore/buffered_channel_size:_1000-4       	   10000	   2414398 ns/op
		BenchmarkSemaphore/buffered_channel_size:_10000-4      	   10000	   2715718 ns/op
	*/
	for _, i := range []int{10, 100, 1000, 10000} {
		b.Run(fmt.Sprintf("buffered channel size: %d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				process(j, i)
			}
		})
	}
}
