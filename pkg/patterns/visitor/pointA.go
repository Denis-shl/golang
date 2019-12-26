package visitor

// PointA ...
type PointA interface {
	NamePointA() string
	Pointer
}

// pointA ...
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

// NewPointA ...
func NewPointA() PointA {
	return &pointA{}
}
