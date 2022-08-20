package medium

func allPathsSourceTarget(graph [][]int) [][]int {
	visited := make(map[int][][]int, len(graph))
	paths := visitNode(0, visited, graph)
	var filteredPaths [][]int
	for _, path := range paths {
		for i := 0; i < len(path); i++ {
			if path[i] == len(graph)-1 {
				filteredPaths = append(filteredPaths, path[:i+1])
				break
			}
		}
	}
	return filteredPaths
}

func visitNode(node int, visited map[int][][]int, graph [][]int) [][]int {
	if len(graph[node]) == 0 {
		v := [][]int{{node}}
		visited[node] = v
		return v
	}
	var currentPaths [][]int
	for _, n := range graph[node] {
		if v, ok := visited[n]; !ok {
			visitedNodes := visitNode(n, visited, graph)
			for _, p := range visitedNodes {
				currentPaths = append(currentPaths, [][]int{append([]int{node}, p...)}...)
			}
		} else {
			for _, p := range v {
				currentPaths = append(currentPaths, [][]int{append([]int{node}, p...)}...)
			}
		}
		visited[node] = currentPaths
	}

	return currentPaths
}
