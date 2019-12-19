package ChainOfRes

import "testing"


func TestChainOfRes(t *testing.T){
	t.Run("testing pattern ChainOfResponsibility", func (t *testing.T){
		b := NewHandlerA(NewHandlerC(NewHandlerB(nil)))
		want := true
		got := b.Treat("C")
		if want != got{
			t.Errorf("error want %v got %v", want, got)
		}
	})
	t.Run("testing pattern ChainOfResponsibility", func (t *testing.T){
		b := NewHandlerA(NewHandlerC(NewHandlerB(nil)))
		want := false
		got := b.Treat("X")
		if want != got{
			t.Errorf("error want %v got %v", want, got)
		}
	})
	t.Run("testing pattern ChainOfResponsibility", func (t *testing.T){
		b := NewHandlerA(NewHandlerC(NewHandlerB(NewHandlerA(NewHandlerB(nil)))))
		want := false
		got := b.Treat("adsdf")
		if want != got{
			t.Errorf("error want %v got %v", want, got)
		}
	})
}