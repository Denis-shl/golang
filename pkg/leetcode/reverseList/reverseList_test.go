package reverselist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	reverseListTest = "Test reverse list"
)

func TestReverseList(t *testing.T) {

	list := NewLister()
	obj := NewHandler()
	head := list
	t.Run(reverseListTest, func(t *testing.T) {
		for i := 1; i < 10; i++ {
			list.SetList(NewLister())
			list = list.getNext()
			list.SetVal(i)
		}
		want := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
		head = obj.ReverseList(head)
		got := head.GetArray(head)
		if !assert.EqualValues(t, want, got) {
			t.Errorf("Error reverse func want %v got %v", want, got)
		}
	})
}
