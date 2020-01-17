package main

import (
	"fmt"
	"github.com/Denis-shl/golang/pkg/patterns/memento"
	"github.com/Denis-shl/golang/pkg/patterns/memento/texteditor"
)

func main() {
	backup := memento.NewMemento()
	editor := texteditor.NewTextEditor(backup)

	editor.MakeChanges(1)
	editor.MakeChanges(2)
	editor.MakeChanges(3)
	editor.SetPreviousCopy()
	editor.SetPreviousCopy()
	fmt.Println(editor.ShowCurrentData())
}
