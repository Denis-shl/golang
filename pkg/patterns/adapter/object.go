package adapter

// ObjA implements system to be adapted.
type ObjA struct {
}

// SpecificRequest implementation.
func (a *ObjA) SpecificRequest() string {
	return "Request"
}
