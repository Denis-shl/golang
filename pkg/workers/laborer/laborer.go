package laborer

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
	completedWork int
	workTime      = rand.Intn(5)
)

// Laborer ...
type Laborer interface {
	Work(workerID int, pipelineCh chan work, resultCh chan workersCompletedWorks, wg *sync.WaitGroup)
}


type laborer struct {
}

// Work ...
func (l *laborer) Work(workerID int, pipelineCh chan work, resultCh chan workersCompletedWorks, wg *sync.WaitGroup) {
	var reportEmployee = make(workersCompletedWorks)
	wg.Done()
	for work := range pipelineCh {
		time.Sleep(time.Second * time.Duration(workTime))
		completedWork += work
	}
	reportEmployee[workerID] = completedWork
	resultCh <- reportEmployee
}

// NewLaborer ...
func NewLaborer() Laborer {
	return &laborer{}
}
