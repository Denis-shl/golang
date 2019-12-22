package mediator

import "fmt"

//Mediator ...
type Mediator interface {
	Reach(x string)string
	NewConcrete(p1, p2 Performer)
}

type concrete struct {
	p1, p2 Performer
}

func (c *concrete) Reach(x string)string {

	if x == "Hello who are you" {
		return c.p2.Speak()
	} else if x == "hello" {
		return c.p1.Speak()
	}
	fmt.Println("not sender")
	return "not sender"
}

func (c *concrete) NewConcrete(p1, p2 Performer) {
	c.p1 = p1
	c.p2 = p2
}

func NewMediator() Mediator {
	return &concrete{}
}
