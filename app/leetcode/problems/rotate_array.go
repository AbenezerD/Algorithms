package problems

func rotate(nums []int, k int) []int {
	return append(nums[len(nums)-k:], nums[:len(nums)-k]...)
}
