package main

import (
	"fmt"
	"github.com/Denis-shl/golang/pkg/patterns/memento"
	"github.com/Denis-shl/golang/pkg/patterns/memento/texteditor"
)

func main() {
	backup := memento.NewMemento()
	editor := texteditor.NewTextEditor(backup)

	editor.Set(1)
	editor.Set(2)
	editor.Set(3)
	editor.SetPreviousCopy()
	editor.SetPreviousCopy()
	fmt.Println(editor.Get())
}
