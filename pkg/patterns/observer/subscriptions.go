package observer

import "fmt"

// ConcreteNotify ...
type ConcreteNotify struct {
	data map[Observer]struct{}
}

// Register ...
func (s *ConcreteNotify) Register(sub Observer) {
	if _, ok := s.data[sub]; ok {
		fmt.Println("already signed")
		return
	}
	s.data[sub] = struct{}{}
}

// Notify ...
func (s *ConcreteNotify) Notify(x int64) {
	for j, _ := range s.data {
		j.Update(x)
	}
}

// Deregister ...
func (s *ConcreteNotify) Deregister(sub Observer) {
	if _, ok := s.data[sub]; ok {
		delete(s.data, sub)
	}
}

// NewConcreteNotify ...
func NewConcreteNotify() Notifier {
	return &ConcreteNotify{make(map[Observer]struct{})}
}
