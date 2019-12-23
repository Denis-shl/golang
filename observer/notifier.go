package observer

// Notifier ...
type Notifier interface {
	Register(Observer)
	Deregister(Observer)
	Notify(int64)
}
