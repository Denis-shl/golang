package factory_method

type ConcreteProductC struct {
	action string
}


func (p *ConcreteProductC) Use() string {
	return p.action
}
