package observer

import "fmt"

// ObserveB ...
type ObserveB struct {
	data int64
}

// GetData ...
func (s *ObserveB) GetData() int64 {
	return s.data
}

// Update ...
func (s *ObserveB) Update(data int64) {
	fmt.Printf("*** Observer %d получен: %d\n", s.data, data)
	s.data = data
}

// NewObserveB ...
func NewObserveB() Observer {
	return &ObserveB{}
}
