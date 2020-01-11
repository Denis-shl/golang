package memento

// Originator ...
type Originator interface {
	SetMemento(Memento)
	BackBackup() int
	SetData(x int)
	SetBackup()
}

// TextEditorA ...
type TextEditorA struct {
	Memento
	data int
}

// SetMemento ...
func (t *TextEditorA) SetMemento(memento Memento) {
	t.Memento = memento
}

// SetData ...
func (t *TextEditorA) SetData(x int) {
	t.data = x
}

// SetBackup ...
func (t *TextEditorA) SetBackup() {
	t.Memento.SetBackup(t.data)
}

// BackBackup ...
func (t *TextEditorA) BackBackup() int {
	return t.Memento.GetBackup()
}

// NewTextEditorA ...
func NewTextEditorA() Originator {
	return &TextEditorA{data: 0}
}
