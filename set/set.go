package main

type set struct {
	data []int
}

//repetition does not include
func (s *set) Add(n ...int) {
	for _, number := range n {
		if s.Contains(number) == false {
			s.data = append(s.data, number)
		}
	}
}

// element search function
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
	setNew := new(set)
	setNew = s

	setNew.Add(w.data[0:]...)
	return setNew
}

//Returns the set obtained by the operation of intersecting it with the specified one.
func (s *set)Intersection(w *set)*set{
	setNew := new(set)
	for _, number := range s.data{
		if w.Contains(number) == true{
			setNew.Add(number)
		}
	}
	return setNew
}
//Returns a set that is different current with the specified.
func (s *set)Difference(w *set)*set{
	setNew := new(set)
	for _, number := range s.data{
		if w.Contains(number) == false{
			setNew.Add(number)
		}
	}
	return setNew
}

func InitSet()*set{
	s := new(set)
	return s
}
