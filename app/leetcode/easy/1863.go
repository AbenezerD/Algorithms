package easy

func subsetXORSum(nums []int) int {
	sums := xor(0, map[int]struct{}{}, []int{}, nums)
	totalSum := 0
	for _, sum := range sums {
		totalSum += sum
	}
	return totalSum
}

func xor(index int, visited map[int]struct{}, sums []int, nums []int) []int {
	if len(nums) == 1 {
		return []int{nums[0]}
	}
	xors := xor(index+1, visited, sums, nums[1:])
	for _, xor := range xors {
		if _, ok := visited[index]; !ok {
			sums = append(sums, nums[0])
			visited[index] = struct{}{}
		}
		if xor != nums[0]^xor {
			sums = append(sums, nums[0]^xor)
			sums = append(sums, xor)
		}
	}
	return sums
}
