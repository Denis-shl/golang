package reverselist

// Handler ...
type Handler interface {
	ReverseList(head Lister) Lister
}

type obj struct {
}

// ReverseList ...
func (o *obj) ReverseList(head Lister) Lister {
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

// NewHandler ...
func NewHandler () Handler{
	return &obj{}
}
