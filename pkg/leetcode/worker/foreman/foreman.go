package foreman

import (
	"fmt"
	"sync"
)

var (
	resultCh   = make(chan workersCompletedWorks, 5)
	pipelineCh = make(chan int, 10)
	wg         = sync.WaitGroup{}
)

type (
	workman               = int
	completedWorks        = int
	workersCompletedWorks = map[workman]completedWorks
	work                  = int
)

type workers interface {
	Work(workerID int, pipelineCh chan work, resultCh chan workersCompletedWorks, wg *sync.WaitGroup)
}

type Foreman interface {
	Start(workersCount, quantityJobs int)
}

type foreman struct {
	workers workers
}

// Start ...
func (f *foreman) Start(workersCount, jobsQuantity int) {
	wg.Add(workersCount + 1)
	go f.transferWork(pipelineCh, jobsQuantity, &wg)
	for workerID := 0; workerID < workersCount; workerID++ {
		go f.workers.Work(workerID, pipelineCh, resultCh, &wg)
	}
	wg.Add(1)
	go f.followWork(workersCount, resultCh, &wg)
	wg.Wait() // ожидаем выполнения followWork
}

func (f *foreman) transferWork(pipelineCh chan work, jobsQuantity int, wg *sync.WaitGroup) {
	defer wg.Done()
	for work := 0; work < jobsQuantity; work++ {
		pipelineCh <- 1
	}
	close(pipelineCh)
}

func (f *foreman) followWork(workersCount int, result chan workersCompletedWorks, w *sync.WaitGroup) {
	defer w.Done()
	summingUp := workersCompletedWorks{}
	for i := 0; i < workersCount; i++ {
		summingUp = <-result
		for nameWorkers, count := range summingUp {
			fmt.Println("the worker", nameWorkers, "Works completed - ", count)
		}
	}
}

// NewForeman  ...
func NewForeman(workers workers) Foreman {
	return &foreman{workers:workers}
}
