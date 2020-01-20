package adapter

type adaptee interface {
	SpecificRequest()
}

// Adapter ...
type Adapter interface {
	Request()
}

type adapter struct {
	adaptee
}

func (a *adapter) Request() {
	a.adaptee.SpecificRequest()

}

// NewAdapter ...
func NewAdapter(adaptee adaptee) Adapter {
	return &adapter{adaptee: adaptee}
}
