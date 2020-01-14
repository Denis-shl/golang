package observer

import "fmt"

// ConcreteNotify ...
type ConcreteNotify struct {
	data map[Observer]struct{}
}

// Register ...
func (c *ConcreteNotify) Register(sub Observer) {
	if _, ok := c.data[sub]; ok {
		fmt.Println("already signed")
		return
	}
	c.data[sub] = struct{}{}
}

// Notify ...
func (c *ConcreteNotify) Notify(x int64) {
	for j, _ := range c.data {
		j.Update(x)
	}
}

// Deregister ...
func (c *ConcreteNotify) Deregister(sub Observer) {
	if _, ok := c.data[sub]; ok {
		delete(c.data, sub)
	}
}

// NewConcreteNotify ...
func NewConcreteNotify() Notifier {
	return &ConcreteNotify{make(map[Observer]struct{})}
}
