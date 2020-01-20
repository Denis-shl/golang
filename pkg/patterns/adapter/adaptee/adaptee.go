package adaptee

import "fmt"

type Adaptee interface {
	SpecificRequest()
}

type adaptee struct {
}

func (a *adaptee) SpecificRequest() {
	fmt.Println("Specific Request")
}

// NewAdaptee ...
func NewAdaptee() Adaptee {
	return &adaptee{}
}
