package adapter

// Describe ...
type Describe interface {
	SpecificRequest() string
}

type companyA struct {
}

// SpecificRequest ...
func (c *companyA) SpecificRequest() string {
	return "Request"
}
