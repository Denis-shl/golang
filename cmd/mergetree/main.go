package main

import (
	"fmt"

	merge "github.com/Denis-shl/golang/pkg/leetcode/binaryTree"
)

func main() {
	mas := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	obj := merge.NewHandler()
	mer := merge.NewCreater()

	creater := mer.SetValue(mas)

	n := obj.MergeTrees(creater, creater)
	fmt.Println(n)
}
