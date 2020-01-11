package intersect

// Traverse ...
type Traverse interface {
	Intersection(nums1 []int, nums2 []int) []int
}

type intersect struct {
}

func contains(n int, s []int) int {
	for i, number := range s {
		if number == n {
			return i
		}
	}
	return -1
}

// Intersection ...
func (i *intersect) Intersection(nums1 []int, nums2 []int) []int {
	var sl []int
	for _, number := range nums1 {
		j := contains(number, nums2)
		if j != -1 {
			sl = append(sl, number)
			nums2 = append(nums2[:j], nums2[j+1:]...)
		}
	}
	return sl
}

// NewIntersect ...
func NewIntersect() Traverse {
	return &intersect{}
}
