package bracket

import (
	"testing"
)

const (
	testNotValid   = "Test not valid"
	testValidOne   = "Test valid one"
	testValidTwo   = "Test valid two"
	testValidThree = "Test valid three"
	testValidFour  = "Test valid four"
)

func TestIsValid(t *testing.T) {
	t.Run(testNotValid, func(t *testing.T) {
		str := "()(()()()"
		got := isValid(str)
		want := false
		if got != want {
			t.Errorf("got %v want  %v", got, want)
		}
	})
	t.Run(testValidOne, func(t *testing.T) {
		str := "()()(){}[]"
		got := isValid(str)
		want := true
		if want != got {
			t.Errorf("got %v want  %v", got, want)
		}
	})
	t.Run(testValidTwo, func(t *testing.T) {
		str := ""
		got := isValid(str)
		want := true
		if want != got {
			t.Errorf("got %v want  %v", got, want)
		}
	})
	t.Run(testValidThree, func(t *testing.T) {
		str := "(("
		got := isValid(str)
		want := false
		if want != got {
			t.Errorf("got %v want  %v", got, want)
		}
	})
	t.Run(testValidFour, func(t *testing.T) {
		str := "))"
		got := isValid(str)
		want := false
		if want != got {
			t.Errorf("got %v want  %v", got, want)
		}
	})
}
