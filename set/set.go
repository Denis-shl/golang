package main

type set struct {
	data []int
}

//Add repetition does not include
func (s *set) Add(n ...int) {
	for _, number := range n {
		if s.Contains(number) == false {
			s.data = append(s.data, number)
		}
	}
}

//Contains element search function
func (s *set) Contains(n int) bool {
	for _, number := range s.data {
		if number == n {
			return true
		}
	}
	return false
}

func (s *set) Deleted(n ...int) {
	for i, number := range n {
		if s.Contains(number) == true {
			s.data = append(s.data[:i], s.data[i+1:]...)
		}
	}
}

func (s *set) Union(w *set) *set {
	setNew := s
	setNew.Add(w.data[0:]...)
	return setNew
}

//Intersection returns the set obtained by the operation of intersecting it with the specified one.
func (s *set)Intersection(w *set)*set{
	setNew := new(set)
	for _, number := range s.data{
		if w.Contains(number) == true{
			setNew.Add(number)
		}
	}
	return setNew
}
//Difference returns a set that is different current with the specified.
func (s *set)Difference(w *set)*set{
	setNew := new(set)
	for _, number := range s.data{
		if w.Contains(number) == false{
			setNew.Add(number)
		}
	}
	return setNew
}

func NewSet()*set{
	s := new(set)
	return s
}
