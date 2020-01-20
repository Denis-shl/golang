package memento

type (
	changes = int
	data    = int
)

type storage struct {
	count  int
	backup []int
}

// memento ...
type Memento interface {
	Push(data changes)
	Pop() int
}

// Push ...
func (s *storage) Push(data changes) {
	s.backup = append(s.backup, data)
	s.count++
}

func (s *storage) Pop() data {
	var previousCopy int

	if s.count == 0 {
		return s.backup[0]
	}
	previousCopy = s.backup[s.count-1]
	s.count -= 1
	return previousCopy

}

// NewMemento ...
func NewMemento() Memento {
	return &storage{}
}
