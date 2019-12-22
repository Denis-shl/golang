package mediator

import "fmt"

type senderB struct {
	mediator Mediator
}

func (s *senderB) Send(x string)string {
	return s.mediator.Reach(x)
}

func (s *senderB) SetMediator(mediator Mediator) {
	s.mediator = mediator
}

func (s *senderB) Speak()string {
	fmt.Println("i senderB")
	return "i senderB"
}

func NewSenderB() Performer {
	return &senderB{}
}
