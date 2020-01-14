package factory

type concreteProductA struct {
}

// DoSomething ...
func (c *concreteProductA) DoSomething() string {
	return "I am a PRODUCT A"
}

// NewEnterprise ...
func NewEnterprise() Enterprise {
	return &concreteFactoryA{}
}
