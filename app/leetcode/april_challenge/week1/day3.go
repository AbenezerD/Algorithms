package week1

func maxSubArray(nums []int) int {
	sum := nums[0]
	max := sum
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		sum += num
		if num > sum {
			sum = num
		}

		if sum > max {
			max = sum
		}
	}
	return max
}
