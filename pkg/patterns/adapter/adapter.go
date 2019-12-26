package adapter

// Target ...
type Target interface {
	Request() string
}

// Adapter ...
type adapter struct {
	*objA
}

// Request is an adaptive method.
func (a *adapter) Request() string {
	return a.SpecificRequest()
}

// NewAdapter ...
func NewAdapter(obj *objA) Target {
	return &adapter{obj}
}
