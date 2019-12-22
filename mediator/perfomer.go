package mediator


type Performer interface {
	Send(x string)string
	Speak()string
	SetMediator(mediator Mediator)
}
