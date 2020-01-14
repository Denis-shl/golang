package addtwonumbers

// Calculater ...
type Calculater interface {
	AddTwoNumbers(l1 Lister, l2 Lister) Lister
	addList(val int, head Lister) Lister
}

type addition struct {
}

// AddTwoNumbers ...
func (a *addition) AddTwoNumbers(l1 Lister, l2 Lister) Lister {
	var (
		newList Lister
		sum     int
	)
	for {
		sum += l1.GetVal() + l2.GetVal()
		if sum >= 10 {
			newList = a.addList(sum-10, newList)
			sum = 1
		} else {
			newList = a.addList(sum, newList)
			sum = 0
		}
		if l1.GetNext() == nil && l2.GetNext() == nil && sum == 0 {
			break
		}
		if l1.GetNext() == nil && l2.GetNext() == nil && sum == 1 {
			newList = a.addList(sum, newList)
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

func (a *addition) addList(val int, head Lister) Lister {
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

// NewCalculater ...
func NewCalculater() Calculater {
	return &addition{}
}
