package adapter

// Purpose ...
type Purpose interface {
	Request() string
}

type adaptr struct {
	Describe
}

// Request ...
func (a *adaptr) Request() string {
	return a.SpecificRequest()
}

// NewPursone ...
func NewPursone(obj Describe) Purpose {
	return &adaptr{obj}
}
