package testing_test

import (
	lasthandler "github.com/Denis-shl/golang/pkg/patterns/chainresponsibility/default"
	"github.com/Denis-shl/golang/pkg/patterns/chainresponsibility/handlera"
	"github.com/Denis-shl/golang/pkg/patterns/chainresponsibility/handlerb"
	"github.com/Denis-shl/golang/pkg/patterns/proxy/apache"
	"github.com/Denis-shl/golang/pkg/patterns/proxy/nginx"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testChainOfResponsibilityOne   = "testing pattern ChainOfResponsibility one"
	testChainOfResponsibilityTwo   = "testing pattern ChainOfResponsibility two"
	testChainOfResponsibilityThree = "testing pattern ChainOfResponsibility three"
)

func TestChainOfRes(t *testing.T) {
	serverApache := apache.NewServer()
	serverNginx := nginx.NewServer()
	t.Run(testChainOfResponsibilityOne, func(t *testing.T) {
		handler := handlera.NewHandlerA(handlerb.NewHandlerB(lasthandler.NewDefault(serverApache, serverNginx), serverNginx), serverApache)
		want := "nginx open source code"
		got := handler.ProcessRequest("nginx")
		if !assert.EqualValues(t, want, got) {
			t.Errorf("error want %v got %v", want, got)
		}
	})
	//t.Run(testChainOfResponsibilityTwo, func(t *testing.T) {
	//	b := NewHandlerA(NewHandlerC(handlerb.NewHandlerB(nil)))
	//	want := false
	//	got := b.Treat("X")
	//	if !assert.EqualValues(t, want, got) {
	//		t.Errorf("error want %v got %v", want, got)
	//	}
	//})
	//t.Run(testChainOfResponsibilityThree, func(t *testing.T) {
	//	b := NewHandlerA(NewHandlerC(handlerb.NewHandlerB(NewHandlerA(handlerb.NewHandlerB(nil)))))
	//	want := false
	//	got := b.Treat("W")
	//	if !assert.EqualValues(t, want, got) {
	//		t.Errorf("error want %v got %v", want, got)
	//	}
	//})
}
