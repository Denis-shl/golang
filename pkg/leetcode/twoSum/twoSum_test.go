package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	obj := NewObj()
	t.Run("make the sums of tails of", func(t *testing.T) {
		target := 8
		got := obj.twoSum([]int{1, 2, 4, 4, 5, 6, 8, 9}, target)
		want := []int{1, 5}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want  %v", got, want)
		}
	})

	t.Run("make the sums of tails of", func(t *testing.T) {
		target := 9
		got := obj.twoSum([]int{1}, target)
		want := []int{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want  %v", got, want)
		}
	})
}
