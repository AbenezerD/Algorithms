package week2

func stringShift(s string, shift [][]int) string {
	totalShift := 0
	for _, s := range shift {
		if s[0] == 0 {
			totalShift += s[1]
		} else {
			totalShift -= s[1]
		}
	}

	totalShift %= len(s)
	if totalShift < 0 {
		totalShift += len(s)
	}
	left := s[totalShift:]
	right:= s[:totalShift]
	return left + right
}
