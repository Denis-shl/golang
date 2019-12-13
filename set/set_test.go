package main

import (
	"reflect"
	"testing"

)

func Test_set(t *testing.T) {

	str := []int{1,2,3,4}
	str2 := []int{3,4,5,6}
	t.Run("make the add testing", func(t *testing.T) {
		got := add(str,1,2,3)
		want := []int{1,2,3,4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make the deleted testing", func(t *testing.T) {
		got := deleted(str2,5, 6)
		want := []int{3,4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make the func Contains", func(t *testing.T) {
		got := contains(1, str2)
		want := -1

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make the func Union", func(t *testing.T) {
		got := union(str,str2)
		want := []int{1,2,3,4,6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make the func Intersection", func(t *testing.T) {
		got := intersection(str, str2)
		want := []int{3,4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}