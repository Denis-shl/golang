package adapter

// Purpose ...
type Purpose interface {
	Request() string
}

type adaptr struct {
	*objA
}

// Request ...
func (a *adaptr) Request() string {
	return a.SpecificRequest()
}

// NewAdapter ...
func NewAdapter(obj *objA) Purpose {
	return &adaptr{obj}
}
