package chainofres

type doctorC struct {
	name string
	next Handler
}

// Treat ...
func (p *doctorC) Treat(name string) bool {
	if name != "C" {
		if p.next != nil {
			return p.next.Treat(name)
		}
		return false
	}
	return true
}

// NewHandlerC ...
func NewHandlerC(next Handler) Handler {
	if next == nil {
		return NewDefaulter()
	}
	return &doctorC{next: next}
}
