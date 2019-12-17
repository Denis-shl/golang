package main

import "fmt"

type Animal interface{
	Use() string
}

type CreateAnimal interface{
	CreateAnimals(actions string)Animal
}

type ConcreteCreater struct{
}


func NewAnimal ()CreateAnimal{
	return &ConcreteCreater{}
}

type CreateAnimalCat struct{
	action string
}

func (p *CreateAnimalCat) Use()string{
	return p.action
}


func (p *ConcreteCreater)CreateAnimal(action string)Animal{
	var animals Animal

	if action == "cat"{
		animals = &CreateAnimalCat{"cat"}
	}
	return animals
}

func main (){

	x := NewAnimal()
	w := []CreateAnimal{
		x.CreateAnimals("cat"),
		x.CreateAnimals("cat"),
	}



	fmt.Println(x)
}