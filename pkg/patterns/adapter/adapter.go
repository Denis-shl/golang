package adapter

// Target ...
type Target interface {
	Request() string
}

// Adapter ...
type Adapter struct {
	*ObjA
}

// Request is an adaptive method.
func (a *Adapter) Request() string {
	return a.SpecificRequest()
}

// NewAdapter ...
func NewAdapter(obj *ObjA) Target {
	return &Adapter{obj}
}
