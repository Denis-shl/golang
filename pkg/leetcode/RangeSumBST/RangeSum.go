package rangesum

// Handler ...
type Handler interface {
	RangeSumBST(root Creater, L int, R int) int
}

type obj struct {
}

// rangeSumBST ...
func (o *obj) RangeSumBST(root Creater, L int, R int) int {
	var sum int
	if root != nil {
		if root.GetVal() <= L && root.GetVal() <= R {
			sum += root.GetVal()
		}
		if root.GetVal() < L {
			sum += o.RangeSumBST(root.GetLeft(), L, R)
		}
		if root.GetVal() < R {
			sum += o.RangeSumBST(root.GetRight(), L, R)
		}
	}
	return sum
}

// NewHandler ...
func NewHandler() Handler {
	return &obj{}
}
