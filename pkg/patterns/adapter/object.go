package adapter

// Describe ...
type Describe interface {
	SpecificRequest() string
}

type objA struct {
}

// SpecificRequest ...
func (o *objA) SpecificRequest() string {
	return "Request"
}
