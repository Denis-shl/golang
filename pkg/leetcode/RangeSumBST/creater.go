package rangesum

import (
	"math/rand"
)

// Creater ...
type Creater interface {
	CreateTree(n []int) Creater
	GetVal() int
	GetLeft() Creater
	GetRight() Creater
	setLeft(left Creater)
	setRight(right Creater)
	setVal(num int)
}

type treeNode struct {
	val   int
	left  Creater
	right Creater
}

// SetValue ...
func (t *treeNode) CreateTree(n []int) Creater {
	newTree := NewCreater()
	head := newTree
	for i, num := range n {
		newTree = head
		if i == 0 {
			newTree.setVal(num)
			continue
		}
		for {
			random := rand.Intn(2)
			if random == 1 {
				if newTree.GetLeft() != nil {
					newTree = newTree.GetLeft()
				} else {
					node := NewCreater()
					node.setVal(num)
					newTree.setLeft(node)
					break
				}
			}
			if random == 0 {
				if newTree.GetRight() != nil {
					newTree = newTree.GetRight()
				} else {
					node := NewCreater()
					node.setVal(num)
					newTree.setRight(node)
					break
				}
			}
		}
	}
	return head
}

// SetLeft ...
func (t *treeNode) setLeft(left Creater) {
	t.left = left
}

// SetRight ...
func (t *treeNode) setRight(right Creater) {
	t.right = right
}

// SetVal ...
func (t *treeNode) setVal(num int) {
	t.val = num
}

// GetRight ...
func (t *treeNode) GetRight() Creater {
	return t.right
}

// GetVal ...
func (t *treeNode) GetVal() int {
	return t.val
}

// GetLeft ...
func (t *treeNode) GetLeft() Creater {
	return t.left
}

// NewCreater ...
func NewCreater() Creater {
	return &treeNode{}
}
