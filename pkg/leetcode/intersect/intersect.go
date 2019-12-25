package intersect

// Contains ...
func Contains(n int, s []int) int {
	for i, number := range s {
		if number == n {
			return i
		}
	}
	return -1
}

// Intersection ...
func Intersect(nums1 []int, nums2 []int) []int {
	var sl []int
	for _, number := range nums1 {
		i := Contains(number, nums2)
		if i != -1 {
			sl = append(sl, number)
			nums2 = append(nums2[:i], nums2[i+1:]...)
		}
	}
	return sl
}
