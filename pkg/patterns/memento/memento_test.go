package memento

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

const (
	testNotNil = "Test designers"
	testBackup = "Test backup"
)

func TestMemento(t *testing.T) {
	compA := NewTextEditorA()
	storage := NewStorage()
	
	t.Run(testNotNil, func(t *testing.T) {
		if !assert.NotNil(t, compA, storage) {
			t.Errorf("error test not nil want not nil got: compA %v, storage %v", compA, storage)
		}
	})

	t.Run(testBackup, func(t *testing.T) {
		compA.SetMemento(storage)
		compA.SetData(12)
		compA.SetBackup()
		compA.SetData(134)
		got := compA.BackBackup()
		want := 12
		if !reflect.DeepEqual(got, want) {
			t.Errorf("test error want %v got %v", want, got)
		}
	})
}
