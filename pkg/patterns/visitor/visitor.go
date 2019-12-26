package visitor

// visitor ...
type Visitor interface {
	VisitPointA(points PointA) string
	VisitPointB(points PointB) string
	VisitPointC(points PointC) string
}

// visit ...
type visit struct {
}

// VisitPointA ...
func (v *visit) VisitPointA(points PointA) string {
	return points.NamePointA()
}

// VisitPointB ...
func (v *visit) VisitPointB(points PointB) string {
	return points.NamePointB()
}

// VisitPointC ...
func (v *visit) VisitPointC(points PointC) string {
	return points.NamePointC()
}

// NewVisitor ...
func NewVisitor() Visitor {
	return &visit{}
}
