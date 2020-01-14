package mergeetree

// Merge ...
type Merge interface {
	MergeTree(t1, t2 Creater) Creater
}

type tree struct {
}

// MergeTree ...
func (o *tree) MergeTree(t1, t2 Creater) Creater {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	var t3 Creater

	t3 = NewCreater()
	t3.SetVal(t2.GetVal() + t1.GetVal())
	if t1.GetLeft() != nil {
		if t2.GetLeft() != nil {
			t3.SetLeft(o.MergeTree(t1.GetLeft(), t1.GetRight()))
		}
		if t2.GetLeft() == nil {
			t3.SetLeft(o.MergeTree(t1.GetLeft(), nil))
		}
	} else {
		t3.SetLeft(o.MergeTree(nil, t2.GetLeft()))
	}
	if t1.GetRight() != nil {
		if t2.GetRight() != nil {
			t3.SetRight(o.MergeTree(t1.GetRight(), t2.GetRight()))
		}
		if t2.GetRight() == nil {
			t3.SetRight(o.MergeTree(t1.GetRight(), nil))
		}
	} else {
		t3.SetRight(o.MergeTree(nil, t2.GetRight()))
	}
	return t3
}

// NewMerge ...
func NewMerge() Merge {
	return &tree{}
}
