package addTwoNumbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testOneSuccess = "test one successfully"
	testTwoSuccess = "test two fail"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Run(testOneSuccess, func(t *testing.T) {
		a := &ListNode{Val: 1, Next: nil}
		b := &ListNode{Val: 9, Next: &ListNode{Val: 9}}
		got := addTwoNumbers(a, b)
		want := &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}

		if !assert.EqualValues(t, want, got) {
			t.Errorf("error func addTwoNumbers want %v got %v", want, got)
		}
	})
	t.Run(testTwoSuccess, func(t *testing.T) {
		a := &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: &ListNode{Val: 1}}}
		b := &ListNode{Val: 9, Next: &ListNode{Val: 9}}
		got := addTwoNumbers(a, b)
		want := &ListNode{Val: 0, Next: &ListNode{Val: 9, Next: &ListNode{Val: 2,}}}

		if !assert.EqualValues(t, want, got) {
			t.Errorf("error func addTwoNumbers want %v got %v", want, got)
		}
	})

}
