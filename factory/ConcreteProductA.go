package factory_method

type ConcreteProductA struct {
	action string
}

func (p *ConcreteProductA) Use() string {
	return p.action
}