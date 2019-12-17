package main

import "fmt"

func main(){
	mas := []int{1,2,3,4,5,6,7,8,9}

	mas2 := mas[0:][:4]
	fmt.Println(mas2)
}