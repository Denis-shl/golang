package observer

import "fmt"

type observeA struct {
	data int64
}

// GetData ...
func (o *observeA) GetData() int64 {
	return o.data
}

// Update ...
func (o *observeA) Update(data int64) {
	fmt.Printf("*** Observer %d получен: %d\n", o.data, data)
	o.data = data
}

// NewObserveA ...
func NewObserveA() Observer {
	return &observeA{}
}
