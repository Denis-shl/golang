package bracket

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	t.Run("test not valid", func(t *testing.T) {
		str := "()(()()()"
		got := isValid(str)
		want := false

		if got != want {
			t.Errorf("got %v want  %v", got, want)
		}
	})

	t.Run("test valis", func(t *testing.T) {
		str := "()()(){}[]"
		got := isValid(str)
		want := true

		if want != got {
			t.Errorf("got %v want  %v", got, want)
		}
	})
	t.Run("test valis", func(t *testing.T) {
		str := ""
		got := isValid(str)
		want := true

		if want != got {
			t.Errorf("got %v want  %v", got, want)
		}
	})
	t.Run("test valis", func(t *testing.T) {
		str := "(("
		got := isValid(str)
		want := false

		if want != got {
			t.Errorf("got %v want  %v", got, want)
		}
	})
	t.Run("test valis", func(t *testing.T) {
		str := "))"
		got := isValid(str)
		want := false

		if want != got {
			t.Errorf("got %v want  %v", got, want)
		}
	})
}
