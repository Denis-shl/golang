package mergeetree

import "math/rand"

// Creater ...
type Creater interface {
	SetValue(n []int) Creater
	GetVal() int
	GetLeft() Creater
	GetRight() Creater
	SetLeft(left Creater)
	SetRight(right Creater)
	SetV(num int)
}

type treeNode struct {
	val   int
	left  Creater
	right Creater
}

// SetValue ...
func (t *treeNode) SetValue(n []int) Creater {
	newTree := NewCreater()
	head := newTree
	for i, num := range n {
		newTree = head
		if i == 0{
			newTree.SetV(num)
			continue
		}
		for {
			random := rand.Intn(2)
			if random == 1 {
				if newTree.GetLeft() != nil {
					newTree = newTree.GetLeft()
				} else {
					node := NewCreater()
					node.SetV(num)
					newTree.SetLeft(node)
					break
				}
			}
			if random == 0 {
				if newTree.GetRight() != nil {
					newTree = newTree.GetRight()
				} else {
					node := NewCreater()
					node.SetV(num)
					newTree.SetRight(node)
					break
				}
			}
		}
	}
	return head
}

// GetLeft ...
func (t *treeNode) GetLeft() Creater {
	return t.left
}

// SetLeft ...
func (t *treeNode) SetLeft(left Creater) {
	t.left = left
}

// SetRight ...
func (t *treeNode) SetRight(right Creater) {
	t.right = right
}

// GetRight ...
func (t *treeNode) GetRight() Creater {
	return t.right
}

// SetV ...
func (t *treeNode) SetV(num int) {
	t.val = num
}

// GetCal ...
func (t *treeNode) GetVal() int {
	return t.val
}

// NewCreater ...
func NewCreater() Creater {
	return &treeNode{}
}
