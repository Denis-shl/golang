package ChainOfRes

type doctorA struct {
	name string
	next Handler
}

// Treat ...
func (p *doctorA) Treat(name string) bool {
	if name != "A" {
		if p.next != nil {
			return p.next.Treat(name)
		} else{
			return false
		}
	}
	return true
}

// NewHandler ...
func NewHandlerA(next Handler) Handler {
	if next == nil{
		return &Default{}
	}
	return &doctorA{next: next}
}
