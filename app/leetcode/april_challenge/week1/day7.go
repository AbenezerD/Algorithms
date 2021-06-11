package week1

func countElements(arr []int) int {
	sumCount := make(map[int]int, len(arr))
	elemMap := make(map[int]struct{}, len(arr))
	for _, elem := range arr {
		sumCount[elem+1] += 1
		elemMap[elem] = struct{}{}
	}

	count := 0
	for key := range elemMap {
		count += sumCount[key]
	}
	return count
}
