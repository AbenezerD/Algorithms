package week1

func moveZeroes(nums []int) []int {
	i := 0
	for _, num := range nums {
		if num != 0 {
			nums[i] = num
			i++
		}
	}
	for ; i < len(nums); i++ {
		nums[i] = 0
	}
	return nums
}
