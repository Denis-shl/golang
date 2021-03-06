package isomorphic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIslsomorphis(t *testing.T) {
	obj := NewIsomorphicer()
	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "foo"
		str2 := "bar"
		want := false
		got := obj.IsIsomorphic(str1, str2)
		if !assert.EqualValues(t, want, got) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "aa"
		str2 := "ab"
		want := false
		got := obj.IsIsomorphic(str1, str2)
		if !assert.EqualValues(t, want, got) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "egg"
		str2 := "sad"
		want := false
		got := obj.IsIsomorphic(str1, str2)
		if !assert.EqualValues(t, want, got) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "dgg"
		str2 := "add"
		want := true
		got := obj.IsIsomorphic(str1, str2)
		if !assert.EqualValues(t, want, got) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "title"
		str2 := "paper"
		want := true
		got := obj.IsIsomorphic(str1, str2)
		if !assert.EqualValues(t, want, got) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "titl"
		str2 := "paper"
		want := false
		got := obj.IsIsomorphic(str1, str2)
		if !assert.EqualValues(t, want, got) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
