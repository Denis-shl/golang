package visitor

// Visitor ...
type Visitor interface {
	VisitPointA(points PointerA) string
	VisitPointB(points PointerB) string
	VisitPointC(points PointerC) string
}

type visit struct {
}

// VisitPointA ...
func (v *visit) VisitPointA(points PointerA) string {
	return points.NamePointA()
}

// VisitPointB ...
func (v *visit) VisitPointB(points PointerB) string {
	return points.NamePointB()
}

// VisitPointC ...
func (v *visit) VisitPointC(points PointerC) string {
	return points.NamePointC()
}

// NewVisitor ...
func NewVisitor() Visitor {
	return &visit{}
}
