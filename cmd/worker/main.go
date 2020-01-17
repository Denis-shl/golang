package main

import (
	"github.com/Denis-shl/golang/pkg/leetcode/worker/foreman"
	"github.com/Denis-shl/golang/pkg/leetcode/worker/workers"
)

var (
	countWorking int = 5
	quantityJobs int = 10
)

func main() {
	newWorkers := workers.NewWorkers()
	newForeman := foreman.NewForeman(newWorkers)
	newForeman.Start(countWorking, quantityJobs)
}
