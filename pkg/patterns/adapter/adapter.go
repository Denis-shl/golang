package adapter

// Targer ...
type Targer interface {
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
func NewAdapter(obj *objA) Targer {
	return &adaptr{obj}
}
