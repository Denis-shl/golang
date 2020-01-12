package memento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testBackup = "Test backup"
)

func TestMemento(t *testing.T) {
	compA := NewTextEditorA()
	storage := NewStorage()

	t.Run(testBackup, func(t *testing.T) {
		compA.SetMemento(storage)
		compA.SetData(12)
		compA.SetBackup()
		compA.SetData(134)
		got := compA.BackBackup()
		want := 12
		if !assert.EqualValues(t, want, got) {
			t.Errorf("test error want %v got %v", want, got)
		}
	})
}
