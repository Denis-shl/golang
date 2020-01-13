package recursive

// Handler ...
type Handler interface {
	NumIslands(grid [][]byte) int
}

type obj struct {
}

// NumIslands ...
func (o *obj) NumIslands(grid [][]byte) int {
	var count int
	for i, n := range grid {
		for j, num := range n {
			if num == 49 {
				count++
				rec(grid, i, j)
			}
		}
	}
	return count
}

func rec(grid [][]byte, i int, j int) {
	if grid[i][j] == 49 {
		grid[i][j] = 48
		if i != len(grid)-1 {
			rec(grid, i+1, j)
		}
		if i != 0 {
			rec(grid, i-1, j)
		}
		if j != len(grid[i])-1 {
			rec(grid, i, j+1)
			grid[i][j] = 0
		}
		if j != 0 {
			rec(grid, i, j-1)
		}
	}
}

//NewHandler ...
func NewHandler() Handler {
	return &obj{}
}
