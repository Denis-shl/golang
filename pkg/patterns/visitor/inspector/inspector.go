package inspector

import "github.com/Denis-shl/golang/pkg/patterns/visitor/shop"

type (
	report = string
)

// Inspector ...
type Inspector interface {
	Inspection(shop.Shop) report
}

type inspector struct {
}

func (i *inspector) Inspection(shop shop.Shop) report {
	return shop.Get()
}

// NewInspector ...
func NewInspector() Inspector {
	return &inspector{}
}
