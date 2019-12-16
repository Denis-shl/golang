package main

import "fmt"

type Animal interface{
	speak()
}

type stAnimal struct{
	Animal
	name string
}

type Cat struct{
	stAnimal
	color string
}

func NewAnimal (){
	return
	
}
