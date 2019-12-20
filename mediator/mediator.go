package main

import "fmt"

type Mediator interface {
	Reach(x string)
	NewConcrete(p1, p2 Performer)
}

type concrete struct {
	p1, p2 Performer
}

func (c *concrete) Reach(x string) {

	if x == "Hello who are you" {
		c.p2.Speak()
	} else if x == "hello" {
		c.p1.Speak()
	} else {
		fmt.Println("not sender")
	}
}

func (c *concrete) NewConcrete(p1, p2 Performer) {
	c.p1 = p1
	c.p1 = p2
}

func NewMediator() Mediator {
	return &concrete{}
}

func main() {
	x := NewMediator() // mediator
	s1 := NewSenderA()
	s2 := NewSenderB()

	s1.SetMediator(x)
	s2.SetMediator(x)
	// определяем компании для посредника
	x.NewConcrete(s1, s2)
	//
	s1.Send("Hello who aru")
}
