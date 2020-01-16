package workers

import (
	"math/rand"
	"sync"
	"time"
)

type (
	worker                = int
	completedWorks        = int
	workersCompletedWorks = map[worker]completedWorks
	work                  = int
)

var (
	workTime = rand.Intn(5)
)

// Workers ...
type Workers interface {
	Work(workerID int, pipelineCh chan work, resultCh chan workersCompletedWorks, wg *sync.WaitGroup)
}

type workers struct {
}

// Work ...
func (w *workers) Work(workerID int, pipelineCh chan work, resultCh chan workersCompletedWorks, wg *sync.WaitGroup) {
	var (
		reportEmployee = make(workersCompletedWorks)
		completedWork  int
	)
	wg.Done()
	for work := range pipelineCh {
		time.Sleep(time.Second * time.Duration(workTime))
		completedWork += work
	}
	reportEmployee[workerID] = completedWork
	resultCh <- reportEmployee
}

// NewLaborer ...
func NewWorkers() Workers {
	return &workers{}
}
