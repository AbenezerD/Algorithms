package week1

func singleNumber(nums []int) int {
	occurrence := make(map[int]int, 0)
	for _, num := range nums {
		occurrence[num] += 1
	}

	for num, count := range occurrence {
		if count == 1 {
			return num
		}
	}

	return 0
}

func singleNumber2(nums []int) int {
	occurrence := make(map[int]struct{}, 0)
	for _, num := range nums {
		if _, ok := occurrence[num]; ok {
			delete(occurrence, num)
		} else {
			occurrence[num] = struct{}{}
		}
	}

	for num := range occurrence {
		return num
	}

	return 0
}

func singleNumber3(nums []int) int {
	n := nums[0]
	for i := 1; i < len(nums); i++ {
		n ^= nums[i]
	}

	return n
}
