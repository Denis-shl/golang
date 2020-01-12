package set

// Setter ...
type Setter interface {
	Add(n ...int)
	Contains(n int) bool
	Deleted(n ...int)
	Union(w Setter) Setter
	Intersection(w Setter) Setter
	Difference(w Setter) Setter
	GetData() []int
}

// set ...
type set struct {
	data []int
}

// Add repetition does not include
func (s *set) Add(n ...int) {
	for _, number := range n {
		if s.Contains(number) == false {
			s.data = append(s.data, number)
		}
	}
}

// Contains element search function
func (s *set) Contains(n int) bool {
	for _, number := range s.data {
		if number == n {
			return true
		}
	}
	return false
}

// Deleted ...
func (s *set) Deleted(n ...int) {
	for i, number := range n {
		if s.Contains(number) == true {
			s.data = append(s.data[:i], s.data[i+1:]...)
		}
	}
}

// Union ...
func (s *set) Union(w Setter) Setter {
	setNew := &set{}
	data := s.GetData()
	setNew.Add(data[0:]...)
	return setNew
}

// Intersection returns the set obtained by the operation of intersecting it with the specified one.
func (s *set) Intersection(w Setter) Setter {
	setNew := &set{}
	for _, number := range s.data {
		if w.Contains(number) == true {
			setNew.Add(number)
		}
	}
	return setNew
}

// Difference returns a set that is different current with the specified.
func (s *set) Difference(w Setter) Setter {
	setNew := &set{}
	for _, number := range s.data {
		if w.Contains(number) == false {
			setNew.Add(number)
		}
	}
	return setNew
}

// GetData ...
func (s *set) GetData() []int {
	return s.data
}

// NewSetter ...
func NewSetter() Setter {
	s := &set{}
	return s
}
