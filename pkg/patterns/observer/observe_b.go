package observer

import "fmt"

type observeB struct {
	data int64
}

// GetData ...
func (o *observeB) GetData() int64 {
	return o.data
}

// Update ...
func (o *observeB) Update(data int64) {
	fmt.Printf("*** Observer %d получен: %d\n", o.data, data)
	o.data = data
}

// NewObserveB ...
func NewObserveB() Observer {
	return &observeB{}
}
