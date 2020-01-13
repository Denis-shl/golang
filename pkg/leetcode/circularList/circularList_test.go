package circularlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testCycleList    = "test Cycle list"
	testNotCycleList = "test not Cycle list"
)

func TestHasCycle(t *testing.T) {
	list := NewLister()
	obj := NewCycler()
	head := list
	t.Run(testCycleList, func(t *testing.T) {
		for i := 0; i < 10; i++ {
			listNew := NewLister()
			list.SetList(listNew)
			list = list.GetNext()
			list.SetVal(i)
			if i == 9 {
				list.SetList(head)
			}
		}
		got := obj.HasCycle(head)
		want := true
		if !assert.EqualValues(t, got, want) {
			t.Errorf("error test got %v  want %v", got, want)
		}
	})
	t.Run(testNotCycleList, func(t *testing.T) {
		for i := 0; i < 10; i++ {
			listNew := NewLister()
			list.SetList(listNew)
			list = list.GetNext()
			list.SetVal(i)
		}
		got := obj.HasCycle(head)
		want := false
		if !assert.EqualValues(t, got, want) {
			t.Errorf("error test got %v  want %v", got, want)
		}
	})
}
