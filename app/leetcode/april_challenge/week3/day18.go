package week3

import "math"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			left, top := math.MaxInt16, math.MaxInt16
			if j > 0 {
				left = grid[i][j-1]
			}
			if i > 0 {
				top = grid[i-1][j]
			}
			min := left
			if top < min {
				min = top
			}
			grid[i][j] += min
		}
	}

	return grid[m-1][n-1]
}
