package set

import (
	"reflect"
	"testing"
)

const (
	testAdd          = "make the add testing"
	testDeleted      = "make the deleted testing"
	testContains     = "make the func Contains"
	testUnion        = "make the func Union"
	testIntersection = "make the func Intersection"
	testDifference   = "make the func difference"
)

func Test_set(t *testing.T) {
	got := NewSet()
	sets := NewSet()
	sets.Add(1, 2, 3)
	t.Run(testAdd, func(t *testing.T) {
		got.Add(1, 2, 3)
		want := []int{1, 2, 3}
		if !reflect.DeepEqual(got.data, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run(testDeleted, func(t *testing.T) {
		got.Deleted(1)
		want := []int{2, 3}
		if !reflect.DeepEqual(got.data, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run(testContains, func(t *testing.T) {
		x := got.Contains(2)
		want := true
		if x != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run(testUnion, func(t *testing.T) {
		s := NewSet()
		s = got.Union(sets)
		want := []int{2, 3, 1}
		if !reflect.DeepEqual(s.data, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run(testIntersection, func(t *testing.T) {
		got = NewSet()
		got.Add(1, 2, 3, 4)
		sets = NewSet()
		sets.Add(3, 4)
		got = got.Intersection(sets)
		want := []int{3, 4}
		if !reflect.DeepEqual(got.data, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run(testDifference, func(t *testing.T) {
		got = NewSet()
		got.Add(1, 2, 3, 4)
		sets = NewSet()
		sets.Add(3, 4)
		got = got.Difference(sets)
		want := []int{1, 2}
		if !reflect.DeepEqual(got.data, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
