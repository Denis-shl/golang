package owner

import (
	"github.com/Denis-shl/golang/pkg/patterns/visitor/shop"
)

type (
	reports = []string
	shops   = []shop.Shop
)

// Owner ...
type Owner interface {
	VisitShops() reports
}

type owner struct {
	shops
}

func (o *owner) VisitShops() reports {
	var reports reports
	for _, store := range o.shops {
		report := store.AcceptInspector()
		reports = append(reports, report)
	}
	return reports
}

// NewOwner ...
func NewOwner(stores ...shop.Shop) Owner {
	var allShop shops
	for _, store := range stores {
		allShop = append(allShop, store)
	}
	return &owner{
		shops: allShop,
	}
}
