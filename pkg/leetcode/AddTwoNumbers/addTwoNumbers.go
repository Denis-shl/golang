package addtwonumbers

// Handler ...
type Handler interface {
	AddTwoNumbers(l1 Lister, l2 Lister) Lister
}

type obj struct {
}

// AddTwoNumbers ...
func (o *obj) AddTwoNumbers(l1 Lister, l2 Lister) Lister {
	var (
		newList Lister
		sum     int
	)
	for {
		sum += l1.GetVal() + l2.GetVal()
		if sum >= 10 {
			newList = addList(sum-10, newList)
			sum = 1
		} else {
			newList = addList(sum, newList)
			sum = 0
		}
		if l1.GetNext() == nil && l2.GetNext() == nil && sum == 0 {
			break
		}
		if l1.GetNext() == nil && l2.GetNext() == nil && sum == 1 {
			newList = addList(sum, newList)
			break
		}
		if l1.GetNext() != nil {
			l1 = l1.GetNext()
		} else {
			l1.SetVal(0)
		}
		if l2.GetNext() != nil {
			l2 = l2.GetNext()
		} else {
			l2.SetVal(0)
		}
	}
	return newList
}

func addList(val int, head Lister) Lister {
	next := head
	if head == nil {
		list := NewLister()
		list.SetVal(val)
		return list
	}
	for {
		if next.GetNext() == nil {
			list := NewLister()
			list.SetVal(val)
			next.SetList(list)
			return head
		}
		next = next.GetNext()
	}
}

//NewHandler ...
func NewHandler() Handler {
	return &obj{}
}
