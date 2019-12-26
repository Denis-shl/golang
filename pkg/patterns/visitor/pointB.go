package visitor

// PointerB ...
type PointerB interface {
	NamePointB() string
	Pointer
}

type pointB struct {
}

// NamePointB ...
func (p *pointB) NamePointB() string {
	return "Point B"
}

// Accept ...
func (p *pointB) Accept(visitor Visitor) string {
	return visitor.VisitPointB(p)
}

// NewPointB ...
func NewPointB() PointerB {
	return &pointB{}
}
