package memento

type (
	changes = int
	data    = int
)

type storage struct {
	backup []int
}

// memento ...
type Memento interface {
	SetNewCopy(data changes)
	GetPreviousCopy() int
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
