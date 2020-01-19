package memento

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testBackup = "Test backup"
)

func TestMemento(t *testing.T) {
	memento := NewMemento()
	memento.Push(1)
	memento.Push(2)
	backup = []int{1,2,3,4,5,6}



	t.Run(testBackup, func(t *testing.T) {
		want := 5
		got := memento.Pop()

		if !assert.EqualValues(t, want, got){
			t.Errorf("%v %v", want, got)
		}
	})
}
