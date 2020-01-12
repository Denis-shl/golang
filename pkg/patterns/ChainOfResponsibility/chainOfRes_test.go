package chainofres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testChainOfResponsibilityOne   = "testing pattern ChainOfResponsibility one"
	testChainOfResponsibilityTwo   = "testing pattern ChainOfResponsibility two"
	testChainOfResponsibilityThree = "testing pattern ChainOfResponsibility three"
)

func TestChainOfRes(t *testing.T) {
	t.Run(testChainOfResponsibilityOne, func(t *testing.T) {
		b := NewHandlerA(NewHandlerC(NewHandlerB(nil)))
		want := true
		got := b.Treat("C")
		if !assert.EqualValues(t, want, got) {
			t.Errorf("error want %v got %v", want, got)
		}
	})
	t.Run(testChainOfResponsibilityTwo, func(t *testing.T) {
		b := NewHandlerA(NewHandlerC(NewHandlerB(nil)))
		want := false
		got := b.Treat("X")
		if !assert.EqualValues(t, want, got) {
			t.Errorf("error want %v got %v", want, got)
		}
	})
	t.Run(testChainOfResponsibilityThree, func(t *testing.T) {
		b := NewHandlerA(NewHandlerC(NewHandlerB(NewHandlerA(NewHandlerB(nil)))))
		want := false
		got := b.Treat("W")
		if !assert.EqualValues(t, want, got) {
			t.Errorf("error want %v got %v", want, got)
		}
	})
}
