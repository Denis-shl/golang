package memento

// Memento ...
type Memento interface {
	SetBackup(int)
	GetBackup() int
}

// storage ...
type storage struct {
	backup []int
}

// SetBackup ...
func (s *storage) SetBackup(x int) {
	s.backup = append(s.backup, x)
}

// GetBackup ...
func (s *storage) GetBackup() int {
	var ret int
	l := len(s.backup)
	if l >= 1 {
		ret = s.backup[l-1]
		s.backup = append(s.backup[:l-1])
	}
	return ret
}

// NewMemento ...
func NewMemento() Memento {
	return &storage{}
}
