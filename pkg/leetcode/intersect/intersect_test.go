package intersect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testIntersect = "Test Intersection"
)

func TestIntersect(t *testing.T) {
	obj := NewIntersect()
	t.Run(testIntersect, func(t *testing.T) {
		var (
			nums1 = []int{1, 2, 2, 1}
			nums2 = []int{2, 2}
		)
		want := []int{2, 2}
		got := obj.Intersection(nums1, nums2)
		if !assert.EqualValues(t, want, got) {
			t.Errorf("error test intersect want %v got %v", want, got)
		}
	})
}
