package rangeSum

import "fmt"

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// rangeSumBST ...
func rangeSumBST(root *TreeNode, L int, R int) int {
	var sum int
	if root != nil {
		if root.Val <= L && root.Val <= R {
			sum += root.Val
		}
		if root.Val < L {
			fmt.Println(root.Val)
			sum += rangeSumBST(root.Left, L, R)
		}
		if root.Val < R {
			fmt.Println(root.Val)
			sum += rangeSumBST(root.Right, L, R)
		}
	}
	return sum
}
