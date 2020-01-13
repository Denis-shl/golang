package addtwonumbers

// Lister ...
type Lister interface {
	SetVal(num int)
	SetList(newList Lister)
	GetArray(head Lister) []int
	GetNext() Lister
	GetVal() int
}

type listNode struct {
	val  int
	next Lister
}

// SetList ...
func (l *listNode) SetList(newList Lister) {
	l.next = newList
}

// SetVal ...
func (l *listNode) SetVal(num int) {
	l.val = num
}

// GetArray func to convert a list to an array
func (l *listNode) GetArray(head Lister) []int {
	var array []int
	for {
		if head == nil {
			return array
		}
		array = append(array, head.GetVal())
		head = head.GetNext()
	}
}

// GetNext ...
func (l *listNode) GetNext() Lister {
	return l.next
}

// GetVal ...
func (l *listNode) GetVal() int {
	return l.val
}

// NewLister ...
func NewLister() Lister {
	return &listNode{}
}
