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
func (s *singleton) Add() {
	s.Lock()
	defer s.Unlock()
	s.count++
}

// Done decrement the count
func (s *singleton) Done() {
	s.Lock()
	defer s.Unlock()
	s.count--
}

// GetValue receiving count
func (s *singleton) GetValue() int {
	return s.count
}

// NewSingleton object creation
func NewSingleton() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
