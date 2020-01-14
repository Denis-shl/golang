package chainofres

import "fmt"

// Defaulter ...
type Defaulter interface{
	Treat(name string) bool
}
type defaul struct {
}

// Treat ...
func (p *defaul) Treat(name string) bool {
	fmt.Println("couldn't find a doctor")
	return false
}

// NewDefaulter ...
func NewDefaulter() Defaulter{
	return &defaul{}
}
