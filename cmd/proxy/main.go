package main

import (
	"fmt"

	"github.com/Denis-shl/golang/pkg/patterns/proxy"
	"github.com/Denis-shl/golang/pkg/patterns/proxy/apache"
	"github.com/Denis-shl/golang/pkg/patterns/proxy/nginx"
)

func main() {
	serverApache := apache.NewServer()
	serverNginx := nginx.NewServer()
	newProxy := proxy.NewProxy(serverApache, serverNginx)
	response := newProxy.Request("nginx")
	fmt.Println(response)
	response = newProxy.Request("apache")
	fmt.Println(response)
	response = newProxy.Request("apache + nginx")
	fmt.Println(response)
	fmt.Println(newProxy.GetRequestHistory())
}
