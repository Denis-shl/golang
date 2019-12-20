package main

type Performer interface {
	Send(x string)
	Speak()
	SetMediator(mediator Mediator)
}
