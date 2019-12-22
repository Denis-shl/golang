package mediator



import "fmt"

type senderA struct {
	mediator Mediator
}

func (s *senderA) Send(x string)string {
	return s.mediator.Reach(x)
}

func (s *senderA) SetMediator(mediator Mediator) {
	s.mediator = mediator
}

func (s *senderA) Speak()string {
	fmt.Println("i senderA")
	return "i senderA"
}

func NewSenderA() Performer {
	return &senderA{}
}
