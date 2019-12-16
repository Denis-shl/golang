package worker

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func work(id int, jobs chan int, res chan map[int]int, w *sync.WaitGroup) {
	var x int
	defer w.Done()
	timeRand := rand.Intn(4)
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

func foreman(working int, result chan map[int]int, w *sync.WaitGroup) {
	defer w.Done()
	for i := 1; i <= working; i++ {
		<-result
	}
}

func workerPool(working int) {
	result := make(chan map[int]int, working) // map[working][work complete]
	jobs := make(chan int, 100)
	w := sync.WaitGroup{}
	for i := 1; i <= working; i++ {
		w.Add(1)
		go work(i, jobs, result, &w)
	}
	go func(jobs chan int) {
		defer close(jobs)

		for i := 0; i < 8; i++ {
			jobs <- 1
		}
	}(jobs)
	w.Add(1)
	go foreman(working, result, &w)
	w.Wait()
}
