package observer

// Observer ...
type Observer interface {
	Update(int64)
	GetData() int64
}
