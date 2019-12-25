package chainOfRes

// Handler ...
type Handler interface {
	Treat(string) bool
}
