package hard

// For example, all the following are valid numbers:
// ["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"],
// while the following are not valid numbers: ["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]
func isNumber(s string) bool {
	digits := map[rune]bool{
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}
	var prev rune
	eAllowed, dotAllowed := true, true
	digitsNeeded := false
	start := true
	for _, c := range s {
		switch c {
		case '+', '-':
			if !start && prev != 'e' && prev != 'E' {
				return false
			}
			digitsNeeded = true
		case '.':
			if !dotAllowed {
				return false
			}
			dotAllowed = false
			if !digits[prev] {
				digitsNeeded = true
			}
		case 'e', 'E':
			if start || digitsNeeded || !eAllowed {
				return false
			}
			eAllowed = false
			dotAllowed = false
			digitsNeeded = true
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			digitsNeeded = false
		default:
			return false

		}
		start = false
		prev = c
	}

	return !digitsNeeded
}
