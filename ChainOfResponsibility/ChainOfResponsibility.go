package chainOfRes

type Handler interface {
	Treat(string) bool
}
