package circularList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (p *ListNode) Append(head *ListNode, listNew *ListNode) *ListNode {
	tmp := &head
	for {
		if (*tmp).Next == nil {
			(*tmp).Next = listNew
			return head
		}
		tmp = &(*tmp).Next
	}
}

func TestHasCycle(t *testing.T) {

	t.Run("test Cycle list", func(t *testing.T) {
		head := &ListNode{}
		for i := 0; i < 10; i++ {
			listNew := &ListNode{}
			head.Append(head, listNew)
		}
		head.Next.Next.Next.Next.Next.Next.Next.Next = head
		got := hasCycle(head)
		want := true
		if !assert.EqualValues(t, got, want) {
			t.Errorf("error test got %v  want %v", got, want)
		}
	})
}
