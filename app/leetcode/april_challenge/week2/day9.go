package week2

func backspaceCompare(S, T string) bool {
	return removeBackspace(S) == removeBackspace(T)
}

func removeBackspace(str string) string {
	arr := make([]uint8, len(str))
	backspaceCount := 0
	j := len(str) - 1
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '#' {
			backspaceCount++
		} else if backspaceCount == 0 {
			arr[j] = str[i]
			j--
		} else {
			backspaceCount--
		}
	}
	return string(arr[j+1:])
}
