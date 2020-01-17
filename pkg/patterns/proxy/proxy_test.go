package proxy

import (
	"github.com/Denis-shl/golang/pkg/patterns/proxy/apache"
	"github.com/Denis-shl/golang/pkg/patterns/proxy/nginx"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testProxyApache      = "testing Proxy - server Apache"
	testProxyNginx       = "testing Proxy - server Nginx"
	testProxyApacheNginx = "testing Proxy - Apache - Server"
	testProxy            = "testing Proxy"
)

func TestProxy(t *testing.T) {
	serverApache := apache.NewServer()
	serverNginx := nginx.NewServer()
	proxy := NewProxy(serverApache, serverNginx)

	t.Run(testProxyApache, func(t *testing.T) {
		want := "apache open source code"
		got := proxy.Request("apache")
		if !assert.EqualValues(t, want, got) {
			t.Errorf("want %v got %v", want, got)
		}
	})
	t.Run(testProxyNginx, func(t *testing.T) {
		want := "nginx open source code"
		got := proxy.Request("nginx")
		if !assert.EqualValues(t, want, got) {
			t.Errorf("want %v got %v", want, got)
		}
	})
	t.Run(testProxyApacheNginx, func(t *testing.T) {
		want := "nginx open source code,apache open source code"
		got := proxy.Request("apache nginx")
		if !assert.EqualValues(t, want, got) {
			t.Errorf("want %v got %v", want, got)
		}
	})
	t.Run(testProxy, func(t *testing.T) {
		want := "400"
		got := proxy.Request("closed source code")
		if !assert.EqualValues(t, want, got) {
			t.Errorf("want %v got %v", want, got)
		}
	})
}
