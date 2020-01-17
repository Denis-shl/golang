package main

import (
	"github.com/Denis-shl/golang/pkg/patterns/visitor/kashyyyk"
	"github.com/Denis-shl/golang/pkg/patterns/visitor/planet"
	"github.com/Denis-shl/golang/pkg/patterns/visitor/yavin"
)

func main (){
	planet.NewPlanet()
	kashyyyk.NewPlanet()
	yavin.NewPlanet()
}
