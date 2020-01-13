package chainofres

type doctorA struct {
	name string
	next Handler
}

// Treat ...
func (p *doctorA) Treat(name string) bool {
	if name != "A" {
		if p.next != nil {
			return p.next.Treat(name)
		}
		return false
	}
	return true
}

// NewHandlerA ...
func NewHandlerA(next Handler) Handler {
	if next == nil {
		return NewDefaulter()
	}
	return &doctorA{next: next}
}