package mediator

import "fmt"

type senderA struct {
	mediator Mediator
}

// Send ...
func (s *senderA) Send(x string) string {
	return s.mediator.Reach(x)
}

// SetMediator ...
func (s *senderA) SetMediator(mediator Mediator) {
	s.mediator = mediator
}

// speak ...
func (s *senderA) speak() string {
	fmt.Println("i senderA")
	return "i senderA"
}

// NewPerformerA ...
func NewPerformerA() Performer {
	return &senderA{}
}
