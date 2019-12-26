package visitor

// PointB ...
type PointB interface {
	NamePointB() string
	Pointer
}

// pointB ...
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
func NewPointB() PointB {
	return &pointB{}
}
