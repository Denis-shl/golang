package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testProxy = "testing Proxy pattern"
)

func TestProxy(t *testing.T) {
	t.Run(testProxy, func(t *testing.T) {
		x := NewProxer()
		want := "string one"
		x.SetName("string one")
		got := x.GetName()
		if !assert.EqualValues(t, want, got) {
			t.Errorf("testing error want %v got %v", want, got)
		}
	})
}
