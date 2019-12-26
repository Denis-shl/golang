package addTwoNumbers

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// addList ...
func addList(val int, head *ListNode) *ListNode {
	next := head
	if head == nil {
		return &ListNode{Val: val}
	}
	for {
		if next.Next == nil {
			next.Next = &ListNode{Val: val}
			return head
		}
		next = next.Next
	}
	return head
}

// addTwoNumbers ...
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		newList *ListNode
		sum     int
	)
	for {
		sum += l1.Val + l2.Val
		if sum >= 10 {
			newList = addList(sum-10, newList)
			sum = 1
		} else {
			newList = addList(sum, newList)
			sum = 0
		}
		if l1.Next == nil && l2.Next == nil && sum == 0 {
			break
		}
		if l1.Next == nil && l2.Next == nil && sum == 1 {
			newList = addList(sum, newList)
			break
		}
		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1.Val = 0
		}
		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2.Val = 0
		}
	}
	return newList
}
