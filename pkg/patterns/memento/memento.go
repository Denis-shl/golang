package memento

type (
	changes = int
	data    = int
)

// memento ...
type Memento interface {
	SetNewCopy(data changes)
	GetPreviousCopy() int
}

// storage ...
type storage struct {
	backup []int
}

// SetNewCopy ...
func (s *storage) SetNewCopy(data changes) {
	s.backup = append(s.backup, data)
}

// GetPreviousCopy ...
func (s *storage) GetPreviousCopy() data {
	var previousCopy int

	backupID := len(s.backup) - 1
	if backupID >= 1 {
		previousCopy = s.backup[backupID-1]
		s.backup = append(s.backup[:backupID])
	}
	return previousCopy

}

// NewMemento ...
func NewMemento() Memento {
	return &storage{}
}
