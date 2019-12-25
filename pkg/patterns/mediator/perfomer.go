package mediator

// Performer ...
type Performer interface {
	Send(x string) string
	speak() string
	SetMediator(mediator Mediator)
}
