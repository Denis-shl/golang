package main

import (
	"fmt"

	creater "github.com/Denis-shl/golang/pkg/leetcode/rangesumbst"
)

func main() {
	mas := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	tree := creater.NewCreater()
	tree = tree.CreateTree(mas)
	obj := creater.NewRangeTree()
	res := obj.RangeSumBST(tree, 4, 10)
	fmt.Println(res)
}
