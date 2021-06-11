package week2

func findMaxLength1(nums []int) int {
	arr := make([]int, 2 * len(nums) + 1)
	for i := range arr {
		arr[i] = -2
	}
	arr[len(nums)] = -1
	maxlen, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count += -1
		} else {
			count += 1
		}
		if arr[count + len(nums)] >= -1 {
			i2 := i - arr[count+len(nums)]
			if i2 > maxlen {
				maxlen = i2
			}
		} else {
			arr[count + len(nums)] = i
		}

	}
	return maxlen
}

func findMaxLength(nums []int) int {
	zeroCount := make(map[int]int, 0)
	zeroCount[0] = -1
	maxCount, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count += -1
		} else {
			count += 1
		}
		if existingCount, ok := zeroCount[count]; ok {
			i2 := i - existingCount
			if i2 > maxCount {
				maxCount = i2
			}
		} else {
			zeroCount[count] = i
		}
	}
	return maxCount
}
