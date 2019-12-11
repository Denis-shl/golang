package main

import "testing"
import "reflect"

func TestTwoSum(t *testing.T) {

	target := 3
	got := twoSum([]int {1,2,3,4,5,6,7,8,9}, target)
	want := []int {0, 1}

	t.Run("make the sums of tails of", func(t *testing.T){
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want  %v", got, want)
	}
	})
}