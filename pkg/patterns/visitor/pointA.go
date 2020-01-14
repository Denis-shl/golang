package visitor

// PointerA ...
type PointerA interface {
	NamePointA() string
	Pointer
}

type pointA struct {
}

// NamePointA ...
func (p *pointA) NamePointA() string {
	return "Point A"
}

// Accept ...
func (p *pointA) Accept(visitor Visitor) string {
	return visitor.VisitPointA(p)
}

// NewPointerA ...
func NewPointerA() PointerA {
	return &pointA{}
}
