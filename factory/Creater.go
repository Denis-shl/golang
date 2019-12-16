package factory_method

type Create interface {
	CreateProduct(action string) Producter // Factory method
}