package visitor

// Pointer ...
type Pointer interface {
	Accept(visitor Visitor) string
}
