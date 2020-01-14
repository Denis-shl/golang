package chainofres

// Handler ...
type Handler interface {
	Treat(string) bool
}
