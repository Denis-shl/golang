package main

import (
	work "github.com/Denis-shl/golang/pkg/leetcode/worker"
)

type worker interface {
	StartJobs()
}

func getObj() worker {
	var (
		CountWorking int = 10
		TimeJobs     int = 5
		QuantityJobs int = 10
	)
	return work.NewWorker(CountWorking, TimeJobs, QuantityJobs)
}

func main() {
	obj := getObj()
	obj.StartJobs()
}