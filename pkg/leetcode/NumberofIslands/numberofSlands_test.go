package recursive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testNumIslandsOne = "test func NumIsland"
	testNumIslandsTwo = "test func NumIsland"
)

func TestNumIslands(t *testing.T) {
	obj := NewHandler()
	t.Run(testNumIslandsOne, func(t *testing.T) {
		grid := [][]byte{
			0: {'1', '1', '0', '0', '0'},
			1: {'1', '0', '0', '0', '0'},
			2: {'0', '0', '1', '0', '0'},
			3: {'0', '0', '0', '1', '1'},
		}
		want := 3
		got := obj.NumIslands(grid)
		if !assert.EqualValues(t, want, got) {
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
		got := obj.NumIslands(grid)
		if !assert.EqualValues(t, want, got) {
			t.Errorf("error test func NumIsland want %v got %v", want, got)
		}
	})
}
