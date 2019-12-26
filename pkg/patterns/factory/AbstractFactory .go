package factory

import "fmt"

type Product interface {
	DoSmth()
}

type Factory interface {
	CreateProduct() Product
}

type concreteProductA struct {
}

func (c *concreteProductA) DoSmth() {
	fmt.Println("I am a PRODUCT A")
}

type concreteFactoryA struct {
}

func (c *concreteFactoryA) CreateProduct() Product{
	return &concreteProductA{}
}

func NewConcreteA ()Factory{
	return &concreteFactoryA{}
}

func main (){
	x := NewConcreteA()

	w := x.CreateProduct()

	w.DoSmth()
}

