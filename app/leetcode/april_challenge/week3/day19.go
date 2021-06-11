package week3

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	pivotIndex := binarySearchForPivot(nums, 0, len(nums))
	search := binarySearchForTree(append(nums[pivotIndex:], nums[:pivotIndex]...), 0, len(nums)-1, target)
	if search == -1 {
		return -1
	}
	return (search + pivotIndex) % len(nums)
}

func binarySearchForTree(nums []int, left, right, target int) int {
	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	}
	if left == right {
		return -1
	}
	if nums[mid] > target {
		return binarySearchForTree(nums, left, mid, target)
	} else {
		return binarySearchForTree(nums, mid+1, right, target)
	}
}

func binarySearchForPivot(nums []int, left, right int) int {
	if left == right {
		return left
	}
	mid := (left + right) / 2
	if nums[mid] < nums[0] {
		return binarySearchForPivot(nums, left, mid)
	} else {
		return binarySearchForPivot(nums, mid+1, right)
	}
}
