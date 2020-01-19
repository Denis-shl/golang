package planet

type firstOrder interface {
	Accept(visitor firstOrder) string
}

// Planet ...
type Planet interface {
}

type planet struct {
	visitor firstOrder
	name    string
}



// Accept ...
func (p *planet) Accept() string {
	return
}

// NewPlanet ...
func NewPlanet(visitor firstOrder) Planet  {
	return &planet {visitor: visitor}
}
