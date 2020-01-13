package reverselist

// Lister ...
type Lister interface {
	SetVal(num int)
	SetList(newList Lister)
	GetArray(head Lister) []int
	getNext() Lister
	getVal() int
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

// GetNext ...
func (l *listNode) getNext() Lister {
	return l.next
}

// GetVal ...
func (l *listNode) getVal() int {
	return l.val
}

// GetArray func to convert a list to an array
func (l *listNode) GetArray(head Lister) []int {
	var array []int
	for {
		if head == nil {
			return array
		}
		array = append(array, head.getVal())
		head = head.getNext()
	}
}

// NewLister ...
func NewLister() Lister {
	return &listNode{}
}
