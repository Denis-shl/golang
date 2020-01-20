package target

type adapter interface {
	Request()
}

// Target ...
type Target interface {
	Request()
}

type target struct {
	adapter
}

func (t *target) Request() {
	t.adapter.Request()
}

// NewTarget ...
func NewTarget(adapter adapter) Target {
	return &target{adapter: adapter}
}
