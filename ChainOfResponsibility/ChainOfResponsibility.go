package ChainOfRes

type Handler interface {
	Treat(string) bool
}
