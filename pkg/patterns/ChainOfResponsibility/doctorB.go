package chainofres

type doctorB struct {
	name string
	next Handler
}

// Treat ...
func (p *doctorB) Treat(name string) bool {
	if name != "B" {
		if p.next != nil {
			return p.next.Treat(name)
		}
		return false
	}
	return true
}

// NewHandlerB ...
func NewHandlerB(next Handler) Handler {
	if next == nil {
		return NewDefaulter()
	}
	return &doctorB{next: next}
}
