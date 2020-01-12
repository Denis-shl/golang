package mergeethree

import "fmt"

// Node ...
type Node interface {
	MergeTrees(t1 *treeNode) Node
}

// Merger ...
type Merger interface {
}

type merge struct {
}

func (m *merge) M(t1, t2 *treeNode) Merger {
	x := mergeTrees(t1, t2)

	x.Val = 10
	fmt.Println(x.Val)
	return x
}

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

// mergeTrees ...
func mergeTrees(t1, t2 *treeNode) *treeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	var t3 treeNode
	t3.Val = t2.Val + t1.Val
	if t1.Left != nil {
		if t2.Left != nil {
			t3.Left = mergeTrees(t1.Left, t1.Right)
		}
		if t2.Left == nil {
			t3.Left = mergeTrees(t1.Left, nil)
		}
	} else {
		t3.Left = mergeTrees(nil, t2.Left)
	}
	if t1.Right != nil {
		if t2.Right != nil {
			t3.Right = mergeTrees(t1.Right, t2.Right)
		}
		if t2.Right == nil {
			t3.Right = mergeTrees(t1.Right, nil)
		}
	} else {
		t3.Right = mergeTrees(nil, t2.Right)
	}
	return &t3
}

// NewQ ...
func NewQ() Q{
	return &merge{}
}

func NewNode () Node{
	return 
}
