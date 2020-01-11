package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	obj := NewObj()
	t.Run("make the sums of tails of", func(t *testing.T) {
		target := 8
		got := obj.twoSum([]int{1, 2, 4, 4, 5, 6, 8, 9}, target)
		want := []int{1, 5}

		if !assert.EqualValues(t, want, got) {
			t.Errorf("got %v want  %v", got, want)
		}
	})

	t.Run("make the sums of tails of", func(t *testing.T) {
		target := 9
		got := obj.twoSum([]int{1}, target)
		want := []int{}

		if !assert.EqualValues(t, want, got) {
			t.Errorf("got %v want  %v", got, want)
		}
	})
}
