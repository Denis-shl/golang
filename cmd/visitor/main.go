package main

import (
	"fmt"

	"github.com/Denis-shl/golang/pkg/patterns/visitor/inspector"
	"github.com/Denis-shl/golang/pkg/patterns/visitor/owner"
	"github.com/Denis-shl/golang/pkg/patterns/visitor/shop"
)

func main() {
	inspectorShop := inspector.NewInspector()
	shopA := shop.NewShop("ShopA", inspectorShop)
	shopB := shop.NewShop("ShopB", inspectorShop)
	shopC := shop.NewShop("ShopC", inspectorShop)
	Owner := owner.NewOwner(shopA, shopB, shopC)
	reports := Owner.VisitShops()

	fmt.Println(reports)
}
