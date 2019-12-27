package adapter

// objA ...
type objA struct {
}

// SpecificRequest ...
func (o *objA) SpecificRequest() string {
	return "Request"
}
