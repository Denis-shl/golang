package worker

import (
	"testing"
)

func BenchmarkSample(b *testing.B) {
	var (
		quantityJobs int = 10
		worker       int = 5
		timeJobs     int = 3
	)
	foreman := NewWorker(worker, timeJobs,quantityJobs)
	b.StartTimer()
	foreman.StartJobs()
	b.StopTimer()
}
