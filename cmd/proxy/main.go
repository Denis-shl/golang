package main

import (
	"github.com/Denis-shl/golang/pkg/patterns/proxy"
	"github.com/Denis-shl/golang/pkg/patterns/proxy/server"

	"fmt"
)

func main() {
	newServer := server.NewServer()
	newProxy := proxy.NewProxy(newServer)
	response := newProxy.Request("nginx")
	fmt.Println(response)
	fmt.Println(newProxy.GetRequestHistory())
}
