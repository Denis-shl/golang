package texteditor

type (
	changes = int
	data    = int
)

type memento interface {
	Push(data changes)
	Pop() (previousCopy data)
}

// TextEditor ...
type TextEditor interface {
	SetPreviousCopy()
	MakeChanges(data changes)
	ShowCurrentData() data
}

type textEditor struct {
	memento
	data int
}

// MakeChanges ...
func (t *textEditor) MakeChanges(data changes) {
	t.memento.Push(data)
	t.data = data
}

// ShowCurrentData ...
func (t *textEditor) ShowCurrentData() data {
	return t.data
}

// SetPreviousCopy ...
func (t *textEditor) SetPreviousCopy() {
	t.data = t.memento.Pop()
}

// NewTextEditor ...
func NewTextEditor(memento memento) TextEditor {
	return &textEditor{memento: memento}
}
