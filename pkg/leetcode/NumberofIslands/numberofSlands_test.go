package recursive

import (
	"reflect"
	"testing"
)

const (
	testNumIslandsOne = "test func NumIsland"
	testNumIslandsTwo = "test func NumIsland"
)

func TestNumIslands(t *testing.T) {
	t.Run(testNumIslandsOne, func(t *testing.T) {
		grid := [][]byte{
			0: {'1', '1', '0', '0', '0'},
			1: {'1', '0', '0', '0', '0'},
			2: {'0', '0', '1', '0', '0'},
			3: {'0', '0', '0', '1', '1'},
		}
		want := 3
		got := numIslands(grid)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("error test func NumIsland want %v got %v", want, got)
		}
	})

	t.Run(testNumIslandsTwo, func(t *testing.T) {
		grid := [][]byte{
			0: {'1', '0', '0', '0', '1'},
			1: {'0', '1', '0', '1', '0'},
			2: {'0', '0', '1', '0', '0'},
			3: {'0', '1', '0', '1', '0'},
			4: {'1', '0', '0', '0', '1'},
			5: {'0', '0', '0', '0', '0'},
			6: {'1', '0', '1', '0', '0'},
		}
		want := 11
		got := numIslands(grid)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("error test func NumIsland want %v got %v", want, got)
		}
	})
}
