package mediator

import "fmt"

type senderB struct {
	mediator Mediator
}

// Send ...
func (s *senderB) Send(x string) string {
	return s.mediator.Reach(x)
}

// SetMediator ...
func (s *senderB) SetMediator(mediator Mediator) {
	s.mediator = mediator
}

func (s *senderB) speak() string {
	fmt.Println("i senderB")
	return "i senderB"
}

// NewSenderB ...
func NewSenderB() Performer {
	return &senderB{}
}
