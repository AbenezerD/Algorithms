package medium

type land struct {
	i int
	j int
}

func maxAreaOfIsland(grid [][]int) int {
	visited := make(map[land]struct{}, 0)

	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if _, ok := visited[land{i, j}]; !ok {
				area := visit(visited, land{i, j}, grid)
				if area > max {
					max = area
				}
			}
		}
	}

	return max
}

func visit(visited map[land]struct{}, l land, grid [][]int) int {
	area := 0
	adjacentLands := getAdjacentLand(l, len(grid), len(grid[0]))
	if _, ok := visited[l]; !ok {
		visited[l] = struct{}{}
		if grid[l.i][l.j] == 1 {
			area++
		} else {
			return 0
		}
	} else {
		return 0
	}

	for _, adjacentLand := range adjacentLands {
		var currentArea int
		currentArea = visit(visited, adjacentLand, grid)
		area += currentArea
	}

	return area
}

func visit2(visited map[land]struct{}, l land, grid [][]int) int {
	area := 0
	adjacentLands := getAdjacentLand(l, len(grid), len(grid[0]))
	if _, ok := visited[l]; !ok {
		visited[l] = struct{}{}
		if grid[l.i][l.j] == 1 {
			area++
		} else {
			return 0
		}
	} else {
		return 0
	}

	for _, adjacentLand := range adjacentLands {
		var currentArea int
		currentArea = visit(visited, adjacentLand, grid)
		area += currentArea
	}

	return area
}

func getAdjacentLand(l land, m, n int) []land {
	i, j := l.i, l.j
	adjacentLands := make([]land, 0)
	if i > 0 {
		adjacentLands = append(adjacentLands, land{i - 1, j})
	}
	if j > 0 {
		adjacentLands = append(adjacentLands, land{i, j - 1})
	}
	if i < m-1 {
		adjacentLands = append(adjacentLands, land{i + 1, j})
	}
	if j < n-1 {
		adjacentLands = append(adjacentLands, land{i, j + 1})
	}
	return adjacentLands
}
