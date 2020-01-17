package firstorder

// Visitor ...
type Visitor interface {
	VisitPlanet(points Tatooine) string
}

type visit struct {
}

// VisitTatooine ...
func (v *visit) VisitTatooine(planet Tatooine) string {
	return planet.NamePlanet()
}

// VisitKashyyyk ...
func (v *visit) VisitKashyyyk(planet Kashyyyk) string {
	return planet.NamePlanet()
}

// VisitYavin ...
func (v *visit) VisitYavin(planet Yavin) string {
	return planet.NamePlanet()
}

// NewVisitor ...
func NewVisitor() Visitor {
	return &visit{}
}
