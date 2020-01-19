package visitor

import "github.com/Denis-shl/golang/pkg/patterns/visitor/planet"

type (
	report string
)

// Visitor ...
type Visitor interface {
	Visit(planet.Planet) report
}

type visit struct {
}

// Visit  ...
func (v *visit) Visit(planet planet.Planet) report {

}

// NewVisitor ...
func NewVisitor() Visitor {
	return &visit{}
}
