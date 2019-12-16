package sort

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T){
	t.Run("make the add testing", func(t *testing.T) {

		nums1 := []int{3,2,1,0,0,0}
		nums2 := []int{6,2,4}

		m := 3
		n := 3
		merge(nums1, m, nums2,n)
		want := []int{1,2,2,3,4,6}
		got := nums1
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}