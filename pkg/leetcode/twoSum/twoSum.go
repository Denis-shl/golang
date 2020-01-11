package twosum

// Summator ...
type Summator interface {
	twoSum(nums []int, target int) []int
}

type obj struct {
}

func (o *obj) twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	l := len(nums)

	if l < 2 {
		return []int{}
	}
	for i, number := range nums {
		hash[number] = i
	}
	for i, number := range nums {
		size := target - number
		if j, ok := hash[size]; ok && j != i {
			return []int{i, j}
		}
	}
	return []int{}
}

// NewObj ...
func NewObj() Summator {
	return &obj{}
}
