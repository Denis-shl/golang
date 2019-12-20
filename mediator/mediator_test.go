package main

import (
	"testing"
)

func TestNewMediator(t *testing.T) {
	t.Run("test new mediator", func(t *testing.T){
		mediator := NewMediator()

		if mediator == nil{
			t.Errorf("TestNewMediator error %v", mediator)
		}
	} )
}

//func TestMediator(t *testing.T) {
//	t.Run("test Mediator")
//}
