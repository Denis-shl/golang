package main

/*O(n^2)
func twoSum(nums []int, target int) []int {

	l := len(nums)

	if l < 2 {
		return []int{}
	}
	for i, number := range nums {
		for j := 0; j < l; j++ {
			if i != j {
				if number + nums[j] == target {
					return []int{i, j}
				}
			}
		}
	}
	return []int{}
}
*/

//O(n)
func twoSum(nums []int, target int) []int {
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
		if j, ok :=  hash[size]; ok  && j != i{
			return []int{i, j}
		}
	}
	return []int{}
}