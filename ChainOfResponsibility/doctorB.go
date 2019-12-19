package ChainOfRes

type doctorB struct {
	name string
	next Handler
}

// Treat ...
func (p *doctorB) Treat(name string) bool {
	if name != "B" {
		if p.next != nil {
			return p.next.Treat(name)
		} else {
			return false
		}
	}
	return true
}

// NewHandler ...
func NewHandlerB(next Handler) Handler {
	if next == nil {
		return &Default{}
	}
	return &doctorB{next: next}
	}
