package main

import (
	"fmt"
	"math/rand"
	"time"
)

func work(id int, jobs chan int, res chan map[int]int) {
	x := 0
	timeRand := rand.Intn(4)
	m := make(map[int]int)
	for j := range jobs {
		time.Sleep(time.Second * time.Duration(timeRand))
		x += j
	}
	m[id] += x
	res <- m

}

func workerPool(working int) {
	result := make(chan map[int]int, working) // map[working][work complete]
	jobs := make(chan int, 100)

	for i := 1; i <= working; i++ {
		go work(i, jobs, result)
	}

	go func (jobs chan int) {
		for i := 0; i < 10; i++ {
			jobs <- 1
		}
		close(jobs)
	}(jobs)

	for i := 0; i < working; i++ {
		x := <-result
		fmt.Println(x)
	}

}

func main() {

	workerPool(4)
}
