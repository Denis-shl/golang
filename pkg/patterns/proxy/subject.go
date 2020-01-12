package proxy

// subject provides an interface for a real subject and its surrogate.
type subject interface {
	SetName(string)
	GetName() string
}