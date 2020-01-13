package isomorphic

// Isomorphicer ...
type Isomorphicer interface {
	IsIsomorphic(s string, t string) bool
}

type isomorphic struct {
}

// IsIsomorphic ...
func (i *isomorphic) IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := make(map[int8]int8)
	hashCheck := make(map[int8]int8)
	for i, num := range s {
		num2 := t[i]
		if hash[int8(num)] == 0 {
			if hashCheck[int8(num2)] == 0 {
				hashCheck[int8(num2)]++
				hash[int8(num)] = int8(num2)
			} else {
				return false
			}
		} else if hash[int8(num)] != int8(num2) {
			return false
		}
	}
	return true
}

// NewIsomorphicer ...
func NewIsomorphicer() Isomorphicer {
	return &isomorphic{}
}
