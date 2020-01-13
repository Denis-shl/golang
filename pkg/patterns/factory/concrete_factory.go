package factory

type concreteFactoryA struct {
}

// CreateProduct ...
func (c *concreteFactoryA) CreateProduct() Producer {
	return &concreteProductA{}
}
