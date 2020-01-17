package texteditor

type (
	changes = int
)

type memento interface {
	SetNewCopy(data changes)
	GetPreviousCopy() int
}

// TextEditor ...
type TextEditor interface {
	SetPreviousCopy()
	MakeChanges(data changes)
	ShowCurrentData() changes
}

type textEditor struct {
	memento
	data int
}

// MakeChanges ...
func (t *textEditor) MakeChanges(data changes) {
	t.memento.SetNewCopy(data)
	t.data = data
}

// ShowCurrentData ...
func (t *textEditor) ShowCurrentData() changes {
	return t.data
}

// SetPreviousCopy ...
func (t *textEditor) SetPreviousCopy() {
	t.data = t.memento.GetPreviousCopy()
}

// NewTextEditor ...
func NewTextEditor(memento memento) TextEditor {
	return &textEditor{memento: memento}
}
