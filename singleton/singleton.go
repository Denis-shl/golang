package main

import (
	"fmt"
	"sync"
)

type Singleton interface {
	Add()
	Done()
	GetValue() int
}

type singleton struct {
	sync.Mutex
	count int
}

var instance *singleton
var once sync.Once

func GetInstance() Singleton {

	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}

func (p *singleton) Add() {
	p.Lock()
	defer p.Unlock()
	p.count++
}

func (p *singleton) Done() {
	p.Lock()
	defer p.Unlock()
	p.count--
}
func (p *singleton) GetValue() int {
	return p.count
}

func main() {

	x := GetInstance()
	x.Add()
	x.Add()
	fmt.Printf("%d\n", x.GetValue())

	y := GetInstance()
	y.Add()
	fmt.Printf("%d", y.GetValue())
}
