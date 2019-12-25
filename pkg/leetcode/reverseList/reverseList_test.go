package reverseList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	reverseListTest = "Test reverse list"
)

func TestReverseList(t *testing.T) {
	list := NewList(10)
	t.Run(reverseListTest, func(t *testing.T) {
		want := []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		got := GetArray(reverseList(list))
		if !assert.EqualValues(t, want, got) {
			t.Errorf("Error reverse func want %v got %v", want, got)
		}
	})
}
