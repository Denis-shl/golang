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

// NewAdapter ...
func NewAdapter(obj Describe) Purpose {
	return &adaptr{obj}
}
