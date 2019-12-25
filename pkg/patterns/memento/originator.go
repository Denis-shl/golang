package memento

// Originator ...
type Originator interface {
	SetMemento(Memento)
	NewMemento()
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

// NewMemento ...
func (t *TextEditorA) NewMemento() {
	t.Memento = &storage{}
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
