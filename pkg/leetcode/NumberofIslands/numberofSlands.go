package recursive

// rec recursive traversal
func rec(grid [][]byte, i int, j int) {
	if grid[i][j] == 1 {
		grid[i][j] = 0
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

// NumIslands ...
func NumIslands(grid [][]byte) int {
	var count int
	for i, n := range grid {
		for j, num := range n {
			if num == 1 {
				count++
				rec(grid, i, j)
			}
		}
	}
	return count
}
