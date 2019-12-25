package circularList

// hasCycle ...
func hasCycle(head *ListNode) bool {
	maps := make(map[*ListNode]int)
	list := head
	for {
		if list == nil {
			break
		} else {
			if maps[list] == 0 {
				maps[list]++
			} else {
				return true
			}
		}
		list = list.Next
	}
	return false
}
