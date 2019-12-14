package main

import (
	"testing"
)

func Test_Worketpool(t *testing.T) {

	workerPool(4)
}

func BenchmarkSample(b *testing.B) {
	b.StartTimer()
	workerPool(4)
	b.StopTimer()
}
