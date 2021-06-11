package problems

func longestPalindrome(s string) string {

	return ""
}

func isPalindrome(s string) bool {
	length := len(s)
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}

	return true
}
