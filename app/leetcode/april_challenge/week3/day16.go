package week3

func checkValidString(s string) bool {
	count, extra := 0, 0
	stack := make([]int32, 0)
	for _, char := range s  {
		switch char {
		case '(':
			count++
			stack = append(stack, char)
		case ')':
			count--
			if count + extra < 0 {
				return false
			}
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case '*':
			extra++
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
