package circularlist

// Cycler ...
type Cycler interface {
	HasCycle(head Lister) bool
}

type list struct {
}

// HasCycle ...
func (l *list) HasCycle(head Lister) bool {
	maps := make(map[Lister]int)
	list := head
	for {
		if list == nil {
			break
		} else {
			if maps[list] == 0 {
				maps[list]++
			} else {
				return true
			}
		}
		list = list.GetNext()
	}
	return false
}

// NewCycler ...
func NewCycler() Cycler {
	return &list{}
}
