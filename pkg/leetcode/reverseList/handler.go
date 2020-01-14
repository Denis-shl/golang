package reverselist

// Reverse ...
type Reverse interface {
	ReverseList(head Lister) Lister
}

type list struct {
}

// ReverseList ...
func (l *list) ReverseList(head Lister) Lister {
	var (
		current Lister
		prev    Lister
	)
	for {
		if head == nil {
			break
		}
		current = head.getNext()
		head.SetList(prev)
		prev = head
		head = current
	}
	return prev
}

// NewReverse ...
func NewReverse () Reverse{
	return &list{}
}
