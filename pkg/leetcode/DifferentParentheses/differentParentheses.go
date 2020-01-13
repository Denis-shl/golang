package diffways

import (
	"strconv"
)

// Handler ...
type Handler interface {
	DiffWaysToCompute(input string) []int
}

type obj struct {
}

func (o *obj) DiffWaysToCompute(input string) []int {
	var (
		res, res1, res2 []int
		str1, str2      []rune
	)
	str := []rune(input)
	for i, num := range input {
		if num == '-' || num == '+' || num == '*' {
			str1 := append(str1, str[:i]...)
			str2 := append(str2, str[i+1:]...)
			res1 = o.DiffWaysToCompute(string(str1))
			res2 = o.DiffWaysToCompute(string(str2))
			for _, n1 := range res1 {
				for _, n2 := range res2 {
					if num == '-' {
						res = append(res, n1-n2)
					} else if num == '+' {
						res = append(res, n1+n2)
					} else if num == '*' {
						res = append(res, n1*n2)
					}
				}
			}
		}
	}
	if res == nil {
		x, _ := strconv.Atoi(input)
		res = append(res, x)
	}
	return res
}

// NewHandler ...
func NewHandler() Handler {
	return &obj{}
}
