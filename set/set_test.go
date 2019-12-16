package main

import (
	"reflect"
	"testing"

)

func Test_set(t *testing.T) {

	got := NewSet()
	sets := NewSet()
	sets.Add(1,2,3)

	t.Run("make the add testing", func(t *testing.T) {

		got.Add(1,2,3)
		want := []int{1,2,3}

		if !reflect.DeepEqual(got.data, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make the deleted testing", func(t *testing.T) {
		got.Deleted(1)
		want := []int{2,3}

		if !reflect.DeepEqual(got.data, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make the func Contains", func(t *testing.T) {
		x := got.Contains(2)
		want := true

		if x != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make the func Union", func(t *testing.T) {

		s := NewSet()
		s = got.Union(sets)
		want := []int{2,3,1}

		if !reflect.DeepEqual(s.data, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	got = NewSet()
	got.Add(1,2,3,4)
	sets = NewSet()
	sets.Add(3,4)

	t.Run("make the func Intersection", func(t *testing.T) {
		got = got.Intersection(sets)
		want := []int{3,4}

		if !reflect.DeepEqual(got.data, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	got = NewSet()
	got.Add(1,2,3,4)
	sets = NewSet()
	sets.Add(3,4)
	t.Run("make the func difference", func(t *testing.T) {
		got = got.Difference(sets)
		want := []int{1,2}

		if !reflect.DeepEqual(got.data, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}