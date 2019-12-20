package main

import "fmt"

type senderB struct {
	mediator Mediator
}

func (s *senderB) Send(x string) {
	s.mediator.Reach(x)
}

func (s *senderB) SetMediator(mediator Mediator) {
	s.mediator = mediator
}

func (s *senderB) Speak() {
	fmt.Println("i senderB")
}

func NewSenderB() Performer {
	return &senderB{}
}
