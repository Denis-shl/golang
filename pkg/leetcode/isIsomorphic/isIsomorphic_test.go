package main

import (
	"testing"
)

func TestIslsomorphis(t *testing.T) {
	obj := NewIsomoprhic()
	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "foo"
		str2 := "bar"
		want := false
		got := obj.isIsomorphic(str1, str2)
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "aa"
		str2 := "ab"
		want := false
		got := obj.isIsomorphic(str1, str2)
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "egg"
		str2 := "sad"
		want := false
		got := obj.isIsomorphic(str1, str2)
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "dgg"
		str2 := "add"
		want := true
		got := obj.isIsomorphic(str1, str2)
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "title"
		str2 := "paper"
		want := true
		got := obj.isIsomorphic(str1, str2)
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "titl"
		str2 := "paper"
		want := false
		got := obj.isIsomorphic(str1, str2)
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
