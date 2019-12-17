package main

import (
	"fmt"
	"sync"
)

type set struct{
	data []int
	mx sync.Mutex
}

func (s *set)Add(num ...int){
	for _, number := range num {
		if s.Contains(number) == false {
			s.data = append(s.data, number)
			
		}
		fmt.Println(number)
	}
}


func (s *set)Contains(originNumber int)bool{
	for _, num := range s.data{
		if num == originNumber{
			return true
		}
	}
	return false
}

func (s *set)Deleted(number ...int){
	for i, num := range number{
		if s.Contains(num) == true{
			s.data = append(s.data[:i], s.data[i + 1:]...)
		}
	}
}

//испрвить название переменных
func (s *set)Union(w ...*set) *set {
	var set2 *set
	for _, x := range w{
		set2.Add(x.data[0:]...)
	}
	return set2
}
//тоже самое 
func (s *set)Intersection(x *set) *set{
	var w *set
	for _, num := range s.data{
		if x.Contains(num) == false{
			w.Add(num)
		}
	}
	return w
}

func (s *set)Defference(x *set) *set{
	var w *set
	for _, num := range s.data{
		if x.Contains(num) == false {
			w.Add(num)
		}
	}

	return w
}

func main (){
	x := new(set)

	x.Add(4,2,3,4)
	x.Add(4,2,3,4)
	fmt.Println(x.data)

	x.Contains(1)

	x.Deleted(4)
	fmt.Println(x.data)

}