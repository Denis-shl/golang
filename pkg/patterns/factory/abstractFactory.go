package factory

type concreteProductA struct {
}

// DoSomething ...
func (c *concreteProductA) DoSomething() string {
	return "I am a PRODUCT A"
}

type concreteFactoryA struct {
}

// CreateProduct ...
func (c *concreteFactoryA) CreateProduct() Producer {
	return &concreteProductA{}
}

// NewConcreteA ...
func NewConcreteA() Enterprise {
	return &concreteFactoryA{}
}
