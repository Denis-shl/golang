package main

import (
	"fmt"

	"github.com/Denis-shl/golang/pkg/patterns/chainresponsibility/handlera"
	"github.com/Denis-shl/golang/pkg/patterns/chainresponsibility/handlerb"
	"github.com/Denis-shl/golang/pkg/patterns/chainresponsibility/lasthandler"
	"github.com/Denis-shl/golang/pkg/patterns/proxy/apache"
	"github.com/Denis-shl/golang/pkg/patterns/proxy/nginx"
)

func main() {
	serverApache := apache.NewServer()
	serverNginx := nginx.NewServer()

	handlerDefault := lasthandler.NewLastHandler(serverApache, serverNginx)
	handlerB := handlerb.NewHandlerB(handlerDefault, serverApache)
	handler := handlera.NewHandlerA(handlerB, serverNginx)

	fmt.Println(handler.ProcessRequest("ng inx"))

}
