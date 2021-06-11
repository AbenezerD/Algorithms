package week3

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++
				flip(grid, i, j)
			}
		}
	}

	return count
}

func flip(grid [][]byte, i, j int) {
	if i >= len(grid) || j >= len(grid[0]) ||
		i < 0 || j < 0 || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	flip(grid, i+1, j)
	flip(grid, i-1, j)
	flip(grid, i, j+1)
	flip(grid, i, j-1)
}
