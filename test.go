package main

import (
	"fmt"
	"sync"
)

/*
// not Mytex

type data struct {
	name string
}
func main (){
	maps := map[string]int{"india" : 10, "Russian" : 30, "USA" : 5}
	sl := make([]int, len(maps))

	for _, number := range maps {
		sl = append(sl, number)
	}

	fmt.Println(len(sl))
	sort.Ints(sl)
	fmt.Println(len(sl))

	for _,num := range sl{
		fmt.Println(num)
	}

}
*/

type MyCounters struct{
	mx sync.Mutex
	n map[string]int
}


func NewCounters() *MyCounters {

	return &MyCounters{
		mx: sync.Mutex{},
		n:  make(map[string]int),
	}
}

func main(){

	var a = []int{1,2,3,4}


	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Printf("%p\n", &a)
	fmt.Println("%p", &a[0])

	a = append(a, 1)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Printf("%p\n", &a)
	fmt.Println("%p", &a[0])
}