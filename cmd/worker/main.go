package main

import worker "github.com/Denis-shl/golang/pkg/leetcode/worker"

func main() {
	var (
		CountWorking int = 10
		TimeJobs     int = 5
		QuantityJobs int = 10
	)
	obj := worker.NewForeman(CountWorking, TimeJobs, QuantityJobs)
	obj.StartJobs()
}
