package chainofres

type doctorC struct {
	name string
	next Handler
}

// Treat ...
func (d *doctorC) Treat(name string) bool {
	if name != "C" {
		if d.next != nil {
			return d.next.Treat(name)
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
