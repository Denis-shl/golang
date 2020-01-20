package texteditor

type (
	changes     = int
	currentData = int
)

type memento interface {
	Push(data changes)
	Pop() int
}

// TextEditor ...
type TextEditor interface {
	SetPreviousCopy()
	Set(changes)
	Get() currentData
}

type textEditor struct {
	memento
	data int
}

// MakeChanges ...
func (t *textEditor) Set(data changes) {
	t.data = data
}

func (t *textEditor) Get() currentData {
	return t.data
}

func (t *textEditor) Make() {

}

// SetPreviousCopy ...
func (t *textEditor) SetPreviousCopy() {
	//t.data = t.memento.Pop()
}

// NewTextEditor ...
func NewTextEditor(memento memento) TextEditor {
	return &textEditor{memento: memento}
}
