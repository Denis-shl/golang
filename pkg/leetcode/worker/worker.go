package worker

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

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

// foreman ...
func foreman(working int, result chan map[int]int, w *sync.WaitGroup) {
	defer w.Done()
	for i := 1; i <= working; i++ {
		<-result
	}
}

// workerPool ...
func workerPool(working int, TimeJobs int, QuantityJobs int) {
	result := make(chan map[int]int, working) // map[working][work complete]
	jobs := make(chan int, QuantityJobs)
	w := sync.WaitGroup{}
	for i := 1; i <= working; i++ {
		w.Add(1)
		go work(i, jobs, result, &w, TimeJobs)
	}
	go func(jobs chan int, QuantityJobs int) {
		defer close(jobs)

		for i := 0; i < QuantityJobs; i++ {
			jobs <- 1
		}
	}(jobs, QuantityJobs)
	w.Add(1)
	go foreman(working, result, &w)
	w.Wait() // ожидаем выполнения foreman
}
