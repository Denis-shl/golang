package main

type info struct{
	symb rune
}

func NewList(symb rune)info{
	return &info{symb}
}

func push(p *info, symb rune){
	for i, x := range p{
		
	}

}
//нужна функция которая кладет в конец списка новую скобку 
//Нужна функция которая проверяет открывающая это скобка или нет
//Нужна функция которая проверяет подходит закрывающая скобка к предыдещей закрывающей 


func sko(str string)bool{
	for i, symb := range str{
		if x(symb) == true {
			push(symb)
		}else{
			if c(symb) == true{

			}
		}
	}
	return true
}
