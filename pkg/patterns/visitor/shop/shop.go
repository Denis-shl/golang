package shop

type (
	report   = string
	name     = string
	nameShop = string
)

type inspector interface {
	Inspection(Shop) report
}

// Shop ...
type Shop interface {
	AcceptInspector() report
	Get() nameShop
}

type shop struct {
	name string
	inspector
}

func (s *shop) AcceptInspector() report {
	return s.inspector.Inspection(s)
}

func (s *shop) Get() nameShop {
	return s.name
}

// NewShop ...
func NewShop(store name, inspector inspector) Shop {
	return &shop{
		name:      store,
		inspector: inspector,
	}
}
