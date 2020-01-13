package singelton

import (
	"sync"
)

// Global variable
var (
	instance *singleton
	once     sync.Once
)

// Singleton is an interface
type Singleton interface {
	Add()
	Done()
	GetValue() int
}

type singleton struct {
	sync.Mutex
	count int
}

// Add increase counter by 1
func (p *singleton) Add() {
	p.Lock()
	defer p.Unlock()
	p.count++
}

// Done decrement the count
func (p *singleton) Done() {
	p.Lock()
	defer p.Unlock()
	p.count--
}

// GetValue receiving count
func (p *singleton) GetValue() int {
	return p.count
}

// NewSingleton object creation
func NewSingleton() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
