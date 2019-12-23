package observer

import "fmt"

// ObserveA ...
type ObserveA struct {
	data int64
}

// GetData ...
func (s *ObserveA) GetData() int64 {
	return s.data
}

// Update ...
func (s *ObserveA) Update(data int64) {
	fmt.Printf("*** Observer %d получен: %d\n", s.data, data)
	s.data = data
}

// NewObserveA ...
func NewObserveA() Observer {
	return &ObserveA{}
}
