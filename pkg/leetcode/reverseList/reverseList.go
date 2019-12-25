package reverseList

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// GetArray func to convert a list to an array
func GetArray(head *ListNode) []int {
	var array []int
	for {
		if head == nil {
			return array
		}
		array = append(array, head.Val)
		head = head.Next
	}
}

// NewList ...
func NewList(size int) *ListNode {
	var head, tmp *ListNode
	head = &ListNode{}
	tmp = head
	for i := 1; i < size; i++ {
		tmp.Val = i
		tmp.Next = &ListNode{}
		tmp = tmp.Next
	}
	return head
}

// reverseList ...
func reverseList(head *ListNode) *ListNode {
	var (
		current *ListNode
		prev    *ListNode
	)
	for {
		if head == nil {
			break
		}
		current = head.Next
		head.Next = prev
		prev = head
		head = current
	}
	return prev
}
