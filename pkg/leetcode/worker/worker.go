package worker

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Worker ...
type Worker interface {
	StartJobs()
	workerPool()
}

type foreman struct {
	working      int
	timeJobs     int
	quantityJobs int
}

// StartJobs ...
func (f *foreman) StartJobs() {
	f.workerPool()
}

func (f *foreman) workerPool() {
	result := make(chan map[int]int, f.working) // map[working][work complete]
	jobs := make(chan int, f.quantityJobs)
	w := sync.WaitGroup{}
	for i := 1; i <= f.working; i++ {
		w.Add(1)
		go work(i, jobs, result, &w, f.timeJobs)
	}
	go func(jobs chan int, QuantityJobs int) {
		defer close(jobs)

		for i := 0; i < QuantityJobs; i++ {
			jobs <- 1
		}
	}(jobs, f.quantityJobs)
	w.Add(1)
	go followWork(f.working, result, &w)
	w.Wait() // ожидаем выполнения foreman
}

// work ...
func work(id int, jobs chan int, res chan map[int]int, w *sync.WaitGroup, TimeJobs int) {
	var x int
	defer w.Done()
	timeRand := rand.Intn(TimeJobs)
	m := make(map[int]int)
	fmt.Println("the worker", id, "began to work")
	for j := range jobs {
		time.Sleep(time.Second * time.Duration(timeRand))
		x += j
	}
	m[id] += x
	res <- m
	fmt.Println("the worker", id, "end to work")
}

func followWork(working int, result chan map[int]int, w *sync.WaitGroup) {
	defer w.Done()
	for i := 1; i <= working; i++ {
		<-result
	}
}

// NewForeman ...
func NewForeman(work, time, quantJobs int) Worker {
	return &foreman{
		working:      work,
		timeJobs:     time,
		quantityJobs: quantJobs,
	}
}
