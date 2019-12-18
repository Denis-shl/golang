package singleton

import (
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
