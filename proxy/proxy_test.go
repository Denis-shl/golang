package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	t.Run("testing Proxy pattern", func (t *testing.T){
		x := Proxy{}
		want := "string one"
		x.PutName("string one")
		got := x.GetName()
		if got != want{
			t.Errorf("testing error want %v got %v", want, got)
		}
	})
}
