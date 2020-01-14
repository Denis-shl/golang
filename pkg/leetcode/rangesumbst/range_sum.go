package rangesum

// RangeTree ...
type RangeTree interface {
	RangeSumBST(root Creater, L int, R int) int
}

type tree struct {
}

// RangeSumBST ...
func (t *tree) RangeSumBST(root Creater, L int, R int) int {
	var sum int
	if root != nil {
		if root.GetVal() <= L && root.GetVal() <= R {
			sum += root.GetVal()
		}
		if root.GetVal() < L {
			sum += t.RangeSumBST(root.GetLeft(), L, R)
		}
		if root.GetVal() < R {
			sum += t.RangeSumBST(root.GetRight(), L, R)
		}
	}
	return sum
}

// NewRangeTree ...
func NewRangeTree() RangeTree {
	return &tree{}
}
