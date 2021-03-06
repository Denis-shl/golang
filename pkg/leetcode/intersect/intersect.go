package intersect

// Traverser ...
type Traverser interface {
	Intersection(nums1 []int, nums2 []int) []int
}

type intersect struct {
}

// Intersection ...
func (i *intersect) Intersection(nums1 []int, nums2 []int) []int {
	var sl []int
	for _, number := range nums1 {
		j := i.contains(number, nums2)
		if j != -1 {
			sl = append(sl, number)
			nums2 = append(nums2[:j], nums2[j+1:]...)
		}
	}
	return sl
}

func (i *intersect) contains(n int, s []int) int {
	for j, number := range s {
		if number == n {
			return j
		}
	}
	return -1
}

// NewTraverser ...
func NewTraverser() Traverser {
	return &intersect{}
}
