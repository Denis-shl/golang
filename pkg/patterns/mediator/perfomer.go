package mediator

// Performer ...
type Performer interface {
	Send(x string) string
	SetMediator(mediator Mediator)
	speak() string
}
