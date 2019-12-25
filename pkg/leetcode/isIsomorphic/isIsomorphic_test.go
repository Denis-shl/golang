package main

import (
	"testing"
)

func TestIslsomorphis(t *testing.T) {
	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "foo"
		str2 := "bar"
		want := false
		got := isIsomorphic(str1, str2)
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "aa"
		str2 := "ab"
		want := false
		got := isIsomorphic(str1, str2)
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "egg"
		str2 := "sad"
		want := false
		got := isIsomorphic(str1, str2)
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "dgg"
		str2 := "add"
		want := true
		got := isIsomorphic(str1, str2)
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "title"
		str2 := "paper"
		want := true
		got := isIsomorphic(str1, str2)
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make test islsomorphis", func(t *testing.T) {
		str1 := "titl"
		str2 := "paper"
		want := false
		got := isIsomorphic(str1, str2)
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	//t.Run("make test islsomorphis", func(t *testing.T) {
	//	str1 := "ab"
	//	str2 := "aa"
	//	want := false
	//	got := isIsomorphic(str1, str2)
	//	if want != got {
	//		t.Errorf("got %v want %v", got, want)
	//	}
	//})
}