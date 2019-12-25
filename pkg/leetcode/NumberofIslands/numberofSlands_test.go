package recursive

import (
	"reflect"
	"testing"
)

const (
	testNumIslands = "test func NumIsland"
)

func TestNumIslands(t *testing.T) {
	t.Run(testNumIslands, func(t *testing.T) {
		grid := [][]byte{
			0: {1, 1, 0, 0, 0},
			1: {1, 0, 0, 0, 0},
			2: {0, 0, 1, 0, 0},
			3: {0, 0, 0, 1, 1},
		}
		want := 3
		got := NumIslands(grid)
		if !  reflect.DeepEqual(want, got) {
			t.Errorf("error test func NumIsland want %v got %v", want, got)
		}
	})
}
