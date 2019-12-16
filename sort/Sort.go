package sort

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {

	for i, num := range nums2{
		nums1[m + i] = num
	}
	sort.Ints(nums1)
}