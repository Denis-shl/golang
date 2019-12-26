package adapter

// objA implements system to be adapted.
type objA struct {
}

// SpecificRequest implementation.
func (o *objA) SpecificRequest() string {
	return "Request"
}
