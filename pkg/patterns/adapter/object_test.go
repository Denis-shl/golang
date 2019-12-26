package adapter

import (
	"reflect"
	"testing"
)

const (
	tesNewAdapter = "Test new Adapter"
)

func TestAdapter(t *testing.T) {
	t.Run(tesNewAdapter, func(t *testing.T) {
		adapter := NewAdapter(&objA{})
		got := adapter.Request()
		want := "Request"
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error test want %v got %v", want, got)
		}
	})
}
