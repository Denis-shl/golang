package visitor

// PointerC ...
type PointerC interface {
	NamePointC() string
	Pointer
}

type pointC struct {
}

// NamePointC ...
func (p *pointC) NamePointC() string {
	return "Point C"
}

// Accept ...
func (p *pointC) Accept(visitor Visitor) string {
	return visitor.VisitPointC(p)
}

// NewPointerC ...
func NewPointerC() PointerC {
	return &pointC{}
}
