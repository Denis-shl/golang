package ChainOfRes

type doctorC struct {
	name string
	next Handler
}

// Treat ...
func (p *doctorC) Treat(name string) bool {
	if name != "C" {
		if p.next != nil {
			return p.next.Treat(name)
		} else {
			return false
		}
	}
	return true
}

// NewHandler ...
func NewHandlerC(next Handler) Handler {
	if next == nil {
		return &Default{}
	}
	return &doctorC{next: next}
}