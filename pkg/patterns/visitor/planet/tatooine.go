package planet

type firstOrder interface {
	Accept(visitor firstOrder) string
}

// PointerA ...
type Planet interface {
	NamePlanet() string
	firstOrder
}

type planet struct {
}

// NamePointA ...
func (p *planet) NamePointA() string {
	return "Point A"
}

// Accept ...
func (p *planet) Accept(visitor Visitor) string {
	return visitor.VisitPointA(p)
}

// NewPlanet ...
func NewPlanet() PointerA {
	return &tatooine{}
}
