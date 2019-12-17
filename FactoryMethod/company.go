package factory_method

import "fmt"

type iCompany interface{
	CreateIceCream()
	GettingNameCompany()string
} 

type Inmarco struct{
	name string
}

type RussianCold struct{
	name string
}

func (p *Inmarco)CreateIceCream (){
	fmt.Println("Inmarco icecream create")
	p.name = "Inmarco"
}

func (p *Inmarco)GettingNameCompany() string{
	return p.name
}

func (p *RussianCold)CreateIceCream (){
	fmt.Println("Russian Cold icecream create")
	p.name = "Russian Cold"
}

func (p *RussianCold)GettingNameCompany() string{
	return p.name
}
