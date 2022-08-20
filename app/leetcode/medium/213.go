package medium

func rob(nums []int) int {
	evenMax := 0
	oddMax := 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			evenMax += nums[i]
		} else {
			oddMax += nums[i]
		}
	}

	if len(nums)%2 != 0 {
		min, _ := order(nums[0], nums[len(nums)-1])
		evenMax -= min
	}

	_, max := order(evenMax, oddMax)
	return max
}

func order(num1, num2 int) (int, int) {
	if num1 < num2 {
		return num1, num2
	}

	return num2, num1
}

func robII(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	d1 := getDistanceArray(nums[:n-1])
	d2 := getDistanceArray(nums[1:])

	return getMax(d1[len(d1)-1], d2[len(d2)-1])
}

func getDistanceArray(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	distance := make([]int, n)
	distance[0] = nums[0]
	distance[1] = getMax(nums[0], nums[1])
	for i := 2; i < n; i++ {
		distance[i] = getMax(distance[i-1], nums[i]+distance[i-2])
	}

	return distance
}

func getMax(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}
