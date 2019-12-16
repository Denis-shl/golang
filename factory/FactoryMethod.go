package factory

// struct
type ConcreteCreater struct {
}

func NewCreater() Create {
	return &ConcreteCreater{}
}

func (p *ConcreteCreater) CreateProduct(action string) Producter {
	var product Producter
	switch action {
	case "A":
		product = &ConcreteProductA{action)
		case "B":
			product = &ConcreteProductB{action}
		case "C":
			product = &ConcreteProductC{action}
		default:
			log.Fatal("Unknown Action")
		}
		return product
	}
