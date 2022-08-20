package medium

func maxArea(height []int) int {
	currentHighest := 0
	nextHighest := 0
	maxArea := 0

	for i := 0; i < len(height); i++ {
		currentArea := getArea(currentHighest, i, height)
		nextArea := getArea(nextHighest, i, height)
		if nextArea >= currentArea && nextArea > maxArea {
			maxArea = nextArea
			currentHighest = nextHighest
			nextHighest = i
		} else if currentArea > maxArea {
			maxArea = currentArea
		}
	}

	return maxArea
}

func getArea(i, j int, height []int) int {
	return (j - i) * min(height[i], height[j])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxArea2(height []int) {
	maxArea := getArea(0, len(height)-1, height)

	left, right := 0, len(height)-1
	for i, j := 1, len(height)-2; i < j; {
		if height[left] < height[right] && height[i] > height[left] {

		}
		area := getArea(i, j, height)
		if area > maxArea {
			maxArea = area

		}

	}
}
