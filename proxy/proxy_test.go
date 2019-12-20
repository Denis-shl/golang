package proxy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProxy(t *testing.T) {
	t.Run("testing Proxy pattern", func(t *testing.T) {
		x := Proxy{}
		want := "string one"
		x.PutName("string one")
		got := x.GetName()
		if !assert.EqualValues(t, want, got) {
			t.Errorf("testing error want %v got %v", want, got)
		}
	})
}
