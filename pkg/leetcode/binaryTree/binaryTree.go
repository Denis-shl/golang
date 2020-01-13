package mergeetree

// Handler ...
type Handler interface {
	MergeTrees(t1, t2 Creater) Creater
}

type obj struct {
}

// MergeTrees ...
func (o *obj) MergeTrees(t1, t2 Creater) Creater {
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
			t3.SetLeft(o.MergeTrees(t1.GetLeft(), t1.GetRight()))
		}
		if t2.GetLeft() == nil {
			t3.SetLeft(o.MergeTrees(t1.GetLeft(), nil))
		}
	} else {
		t3.SetLeft(o.MergeTrees(nil, t2.GetLeft()))
	}
	if t1.GetRight() != nil {
		if t2.GetRight() != nil {
			t3.SetRight(o.MergeTrees(t1.GetRight(), t2.GetRight()))
		}
		if t2.GetRight() == nil {
			t3.SetRight(o.MergeTrees(t1.GetRight(), nil))
		}
	} else {
		t3.SetRight(o.MergeTrees(nil, t2.GetRight()))
	}
	return t3
}

// NewHandler ...
func NewHandler() Handler {
	return &obj{}
}
