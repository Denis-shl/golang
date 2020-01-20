package main

import (
	"github.com/Denis-shl/golang/pkg/patterns/adapter/adaptee"
	"github.com/Denis-shl/golang/pkg/patterns/adapter/adapter"
	"github.com/Denis-shl/golang/pkg/patterns/adapter/client"
	"github.com/Denis-shl/golang/pkg/patterns/adapter/target"
)

func main() {
	newAdaptee := adaptee.NewAdaptee()
	newAdapter := adapter.NewAdapter(newAdaptee)
	newTarget := target.NewTarget(newAdapter)
	newClient := client.NewClient()
	newClient.Request(newTarget)
}
