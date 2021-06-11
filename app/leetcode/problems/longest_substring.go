package problems

func lengthOfLongestSubstring(s string) int {
	currentResetIndex := -1
	charIndex := make(map[rune]int, 0)
	currentMax := 0
	count := 0

	for i, char := range s {
		if index, ok := charIndex[char]; ok {
			if index > currentResetIndex {
				currentResetIndex = index
			}
		}
		count = i - currentResetIndex
		if count > currentMax {
			currentMax = count
		}
		charIndex[char] = i
	}

	return currentMax
}
