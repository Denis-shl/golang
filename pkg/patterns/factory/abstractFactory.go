package factory

type concreteProductA struct {
}

// DoSomething ...
func (c *concreteProductA) DoSomething() string {
	return "I am a PRODUCT A"
}

// NewConcrete ...
func NewConcrete() Enterprise {
	return &concreteFactoryA{}
}
