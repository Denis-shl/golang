package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	tesNewAdapter = "Test new Adapter"
)

func TestAdapter(t *testing.T) {
	t.Run(tesNewAdapter, func(t *testing.T) {
		adapter := NewAdapter(&objA{})
		got := adapter.Request()
		want := "Request"
		if !assert.EqualValues(t, got, want) {
			t.Errorf("Error test want %v got %v", want, got)
		}
	})
}
