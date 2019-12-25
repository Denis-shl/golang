package worker

import (
	"testing"
)

func BenchmarkSample(b *testing.B) {
	QuantityJobs := 10
	worker := 5
	TimeJobs := 3

	b.StartTimer()
	workerPool(worker, TimeJobs, QuantityJobs)
	b.StopTimer()
}
