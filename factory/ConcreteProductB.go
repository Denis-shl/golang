package factory_method

type ConcreteProductB struct {
	action string
}

func (p *ConcreteProductB) Use() string {
	return p.action
}