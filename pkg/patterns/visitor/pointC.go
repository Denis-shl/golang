package visitor

// PointC ...
type PointC interface {
	NamePointC() string
	Pointer
}

// pointC ...
type pointC struct {
}

// NamePoint ...
func (p *pointC) NamePointC() string {
	return "Point C"
}

// Accept ...
func (p *pointC) Accept(visitor Visitor) string {
	return visitor.VisitPointC(p)
}

// NewPointC ...
func NewPointC() PointC {
	return &pointC{}
}
