package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testAbstractFactory = "Testing pattern abstract factory"
)

func TestAbstractFactory(t *testing.T) {
	t.Run(testAbstractFactory, func(t *testing.T) {
		x := NewConcrete()
		w := x.CreateProduct()
		got := w.DoSomething()
		want := "I am a PRODUCT A"
		if !assert.EqualValues(t, want, got) {
			t.Errorf("error test want %v got %v", want, got)
		}
	})
}
