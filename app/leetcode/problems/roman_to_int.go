package problems

var romanToIntMap = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	num := 0
	n := len(s)
	for i := 0; i < n; {
		currentChar := s[i]
		if i == n-1 {
			num += romanToIntMap[currentChar]
			return num
		}
		nextChar := s[i+1]
		newVal, skipNext := getVal(currentChar, nextChar)
		num += newVal
		if skipNext {
			i += 2
		} else {
			i += 1
		}
	}

	return num
}

func getVal(currentChar, nextChar uint8) (int, bool) {
	currentVal := romanToIntMap[currentChar]
	nextVal := romanToIntMap[nextChar]

	if currentVal*5 == nextVal || currentVal*10 == nextVal {
		return nextVal - currentVal, true
	}

	return currentVal, false
}
