package mergeetree

import "math/rand"

// Q ...
type Q interface {
	SetValue(n []int) Node
	getLeft()
	getRight()
}

type three struct {
}

// SetValue ...
func SetValue(n []int) Node {
	newTree := *treeNode

	for i, num := range n {
		for {
			if i == 0 {
				newThree = newThree(num)
			}
			x := rand.Intn(1)
			if x == 1 {
				left := newTree.getLeft()
				if left == nil{
					newTree.SetLeft(newtreeNode(num))
				}
			} else {
				newThree.getRight()
			}
			break
		}
	}
	return newThree
}

func (t *treeNode) getLeft() *treeNode {
	return t.Left
}

func (t *treeNode) getRight() *treeNode {
	return t.Right
}

func newtreeNode(num int) *treeNode {
	return &treeNode{Val: num}
}

// NewQ ...
func NewQ() Node {
	return &treeNode{}
}
