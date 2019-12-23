package mediator

import "fmt"

//Mediator ...
type Mediator interface {
	Reach(x string) string
	SetConcrete(p1, p2 Performer)
}

type concrete struct {
	p1, p2 Performer
}

// Reach ...
func (c *concrete) Reach(x string) string {
	if x == "Hello who are you" {
		return c.p2.speak()
	}
	if x == "hello" {
		return c.p1.speak()
	}
	fmt.Println("not sender")
	return "not sender"
}

// SetConcrete
func (c *concrete) SetConcrete(p1, p2 Performer) {
	c.p1 = p1
	c.p2 = p2
}

// NewMediator
func NewMediator() Mediator {
	return &concrete{}
}
