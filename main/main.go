package main

import (
	"fmt"
	ChainOfRes "github.com/Denis-shl/golang/ChainOfResponsibility"
)

func main() {

	b := ChainOfRes.NewHandlerA(ChainOfRes.NewHandlerC(ChainOfRes.NewHandlerB(nil)))
	fmt.Println(b.Treat("C"))
}


