package diffways

import (
	"strconv"
)

// Calculater ...
type Calculater interface {
	DiffWaysToCompute(input string) []int
}

type compute struct {
}

// DiffWaysToCompute ...
func (c *compute) DiffWaysToCompute(input string) []int {
	var (
		res, res1, res2 []int
		str1, str2      []rune
	)
	str := []rune(input)
	for i, num := range input {
		if num == '-' || num == '+' || num == '*' {
			str1 := append(str1, str[:i]...)
			str2 := append(str2, str[i+1:]...)
			res1 = c.DiffWaysToCompute(string(str1))
			res2 = c.DiffWaysToCompute(string(str2))
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

// NewCalculater ...
func NewCalculater() Calculater {
	return &compute{}
}
