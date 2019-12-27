package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testDesigner        = "Testing constructions"
	testAbstractFactory = "Testing pattern abstract factory"
)

func TestAbstractFactory(t *testing.T) {
	t.Run(testDesigner, func(t *testing.T) {
		got := NewConcreteA()
		if !assert.NotNil(t, got) {
			t.Errorf("error test %v", got)
		}
	})
	t.Run(testAbstractFactory, func(t *testing.T) {
		x := NewConcreteA()
		w := x.CreateProduct()
		got := w.DoSomething()
		want := "I am a PRODUCT A"
		if !assert.EqualValues(t, want, got) {
			t.Errorf("error test want %v got %v", want, got)
		}
	})
}
