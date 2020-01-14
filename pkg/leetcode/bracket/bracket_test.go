package bracket

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testNotValid   = "Test not valid"
	testValidOne   = "Test valid one"
	testValidTwo   = "Test valid two"
	testValidThree = "Test valid three"
	testValidFour  = "Test valid four"
)

func TestIsValid(t *testing.T) {
	obj := NewValidater()

	t.Run(testNotValid, func(t *testing.T) {
		str := "()(()()()"
		got := obj.IsValid(str)
		want := false
		if !assert.EqualValues(t, got, want) {
			t.Errorf("got %v want  %v", got, want)
		}
	})
	t.Run(testValidOne, func(t *testing.T) {
		str := "()()(){}[]"
		got := obj.IsValid(str)
		want := true
		if !assert.EqualValues(t, got, want) {
			t.Errorf("got %v want  %v", got, want)
		}
	})
	t.Run(testValidTwo, func(t *testing.T) {
		str := ""
		got := obj.IsValid(str)
		want := true
		if !assert.EqualValues(t, got, want) {
			t.Errorf("got %v want  %v", got, want)
		}
	})
	t.Run(testValidThree, func(t *testing.T) {
		str := "(("
		got := obj.IsValid(str)
		want := false
		if !assert.EqualValues(t, got, want) {
			t.Errorf("got %v want  %v", got, want)
		}
	})
	t.Run(testValidFour, func(t *testing.T) {
		str := "))"
		got := obj.IsValid(str)
		want := false
		if !assert.EqualValues(t, got, want) {
			t.Errorf("got %v want  %v", got, want)
		}
	})
}
