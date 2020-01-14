package circularlist

// Lister ...
type Lister interface {
	SetVal(num int)
	SetList(newList Lister)
	GetNext() Lister
}

type listNode struct {
	next Lister
	val  int
}

// SetList ...
func (l *listNode) SetList(newList Lister) {
	l.next = newList
}

// SetVal ...
func (l *listNode) SetVal(num int) {
	l.val = num
}

// GetNext ...
func (l *listNode) GetNext() Lister {
	return l.next
}

// NewLister ...
func NewLister() Lister {
	return &listNode{}
}
