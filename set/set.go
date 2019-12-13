package main

//repetition does not include
func add(set []int, n ...int)[]int{

	for  _, number := range  n{
		if contains(number, set) == -1 {
			set = append(set, number)
		}
	}
	return set
}

// element search function
func contains (number int, set []int)int{

	for i, nums := range set{
		if nums == number {
			return i
		}
	}
	return -1
}

func deleted(set []int, el ...int) []int{

	for _, number := range el{
		x := contains(number, set)
		if x != -1 {
			set = append(set[:x], set[x + 1:]...)
		}
	}
	return set
}

func union(set ...[]int)[]int{
	var set2  []int

	for _, s := range set{
		set2 = add(set2, s[0:]...)
	}
	return set2
}

//Returns the set obtained by the operation of intersecting it with the specified one.
func intersection(setOne []int, setTwo []int)[]int{
	var s []int

	for _, numb := range setOne{
		if contains(numb, setTwo) != -1 {
			s = add(s, numb)
		}
	}
	return s
}
//Returns a set that is different current with the specified.
func difference(setOne []int, setTwo []int)[]int{
	var set []int

	for _,num := range setOne{
		if contains(num, setTwo) == -1 {
			set = add(set, num)
		}
	}
	return set
}