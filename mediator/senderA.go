package main

import "fmt"

type senderA struct {
	mediator Mediator
}

func (s *senderA) Send(x string) {
	s.mediator.Reach(x)
}

func (s *senderA) SetMediator(mediator Mediator) {
	s.mediator = mediator
}

func (s *senderA) Speak() {
	fmt.Println("i senderA")
}

func NewSenderA() Performer {
	return &senderA{}
}
