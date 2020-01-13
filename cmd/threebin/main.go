package mergeetree

import (
	three "github.com/Denis-shl/golang/pkg/leetcode/binaryTree"
)

func main (){
	val := []int{1,2,3,4,5,6,7,8,9,10}
	x := three.SetValue(val)

	x.MergeTrees()
}