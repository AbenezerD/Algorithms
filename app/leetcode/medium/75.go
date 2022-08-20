package medium

func sortColors(nums []int) []int {
	_0sIndex := -1
	for i, j := 0, len(nums)-1; i <= j; {
		if nums[i] == 0 && _0sIndex != -1 {
			temp := nums[i]
			nums[i] = nums[_0sIndex]
			nums[_0sIndex] = temp
		}
		if nums[i] > nums[j] {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
		} else if nums[i] == 1 && _0sIndex == -1 {
			if _0sIndex == -1 {
				_0sIndex = i
			} else {
				_0sIndex++
			}
		} else if nums[i] != 2 {
			i++
		}
		if nums[j] == 2 {
			j--
		}
	}

	return nums
}

func sortColors2(nums []int) []int {
	_0sIndex := -1
	for i, j := 0, len(nums)-1; i <= j; {
		if nums[i] == 0 && _0sIndex != -1 {
			temp := nums[i]
			nums[i] = nums[_0sIndex]
			nums[_0sIndex] = temp
			_0sIndex++
		}
		if nums[i] > nums[j] {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
		} else if nums[i] == 1 {
			if _0sIndex == -1 {
				_0sIndex = i
			}
			i++
		} else if nums[i] == 0 {
			i++
		}
		if nums[j] == 2 {
			j--
		}
	}

	return nums
}
