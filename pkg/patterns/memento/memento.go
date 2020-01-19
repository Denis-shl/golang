package memento

//var (
//	backup []int
//)

type (
	changes = int
	data    = int
)

type storage struct {
	backup []int
}

// Memento ...
type Memento interface {
	Push(data changes)
	Pop() (previousCopy data)
}

func (s *storage) Push(data changes) {
	s.backup = append(s.backup, data)
}

func (s *storage) Pop() (previousCopy data) {
	backupID := len(s.backup) - 1
	if backupID >= 1 {
		previousCopy = s.backup[backupID-1]
		s.backup = append(s.backup[:backupID])
	}
	return
}

// NewMemento ...
func NewMemento() Memento {
	return &storage{}
}
