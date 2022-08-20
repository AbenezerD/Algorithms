package medium

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	strArray := make([]string, numRows)

	for i := 0; i < len(s); i++ {
		current := i % ((numRows - 1) * 2)
		if current <= numRows-1 {
			strArray[current] += string(s[i])
		} else {
			index := numRows - 1 - current%(numRows-1)
			strArray[index] += string(s[i])
		}
	}

	return strings.Join(strArray, "")
}
