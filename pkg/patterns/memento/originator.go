package memento

// Originator ...
type Originator interface {
	SetMemento(Memento)
	BackBackup() int
	SetData(x int)
	SetBackup()
}

type textEditorA struct {
	Memento
	data int
}

// SetMemento ...
func (t *textEditorA) SetMemento(memento Memento) {
	t.Memento = memento
}

// SetData ...
func (t *textEditorA) SetData(x int) {
	t.data = x
}

// SetBackup ...
func (t *textEditorA) SetBackup() {
	t.Memento.SetBackup(t.data)
}

// BackBackup ...
func (t *textEditorA) BackBackup() int {
	return t.Memento.GetBackup()
}

// NewOriginator ...
func NewOriginator() Originator {
	return &textEditorA{data: 0}
}
