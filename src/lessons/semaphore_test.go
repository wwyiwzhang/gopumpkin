package main

import "testing"

func BenchmarkSemaphore(b *testing.B) {
	for _, i := range []int{1, 10, 100} {
		b.StopTimer()
		b.StartTimer()
		for j := 0; j < b.N; j++ {
			process(j, i)
		}
		b.StopTimer()
		b.ReportAllocs()
	}
}
