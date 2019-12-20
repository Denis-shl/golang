package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	var t3 TreeNode
	t3.Val = t2.Val + t1.Val
	if t1.Left != nil {
		if t2.Left != nil {
			t3.Left = mergeTrees(t1.Left, t2.Left)
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