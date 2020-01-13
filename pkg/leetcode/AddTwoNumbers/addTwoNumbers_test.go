package addtwonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testOneSuccess = "test one successfully"
	testTwoSuccess = "test two fail"
)

func TestAddTwoNumbers(t *testing.T) {
	obj := NewHandler()
	t.Run(testOneSuccess, func(t *testing.T) {
		a := NewLister()
		a.SetVal(1)
		b := NewLister()
		b.SetVal(9)
		b.SetList(NewLister())
		b.GetNext().SetVal(9)

		got := obj.AddTwoNumbers(a, b)
		want := &listNode{val: 0, next: &listNode{val: 0, next: &listNode{val: 1}}}
		if !assert.EqualValues(t, want, got) {
			t.Errorf("error func addTwoNumbers want %v got %v", want, got)
		}
	})
	t.Run(testTwoSuccess, func(t *testing.T) {
		a := &listNode{val: 1, next: &listNode{val: 9, next: &listNode{val: 1}}}
		b := &listNode{val: 9, next: &listNode{val: 9}}
		got := obj.AddTwoNumbers(a, b)
		want := &listNode{val: 0, next: &listNode{val: 9, next: &listNode{val: 2}}}
		if !assert.EqualValues(t, want, got) {
			t.Errorf("error func addTwoNumbers want %v got %v", want, got)
		}
	})

}
