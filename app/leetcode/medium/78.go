package medium

func subsets(nums []int) [][]int {
	s := [][]int{{}}
	for i, num := range nums {
		currentArr := []int{num}
		sub := subsets(nums[i+1:])
		for _, subInSub := range sub {
			s = append(s, append(currentArr, subInSub...))
		}
	}

	return s
}
