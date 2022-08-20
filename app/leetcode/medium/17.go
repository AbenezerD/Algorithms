package medium

import "strconv"

var digitToChar = map[int][]rune{
	2: {'a', 'b', 'c'},
	3: {'d', 'e', 'f'},
	4: {'g', 'h', 'i'},
	5: {'j', 'k', 'l'},
	6: {'m', 'n', 'o'},
	7: {'p', 'q', 'r', 's'},
	8: {'t', 'u', 'v'},
	9: {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	groups := make([][]rune, len(digits))
	for i, c := range digits {
		digit, _ := strconv.Atoi(string(c))
		groups[i] = digitToChar[digit]
	}

	return traverse17(groups)
}

func traverse17(children [][]rune) []string {
	if len(children) == 1 {
		strArr := make([]string, len(children[0]))
		for i, char := range children[0] {
			strArr[i] = string(char)
		}
		return strArr
	}

	var traversed []string
	childTraversed := traverse17(children[1:])
	for _, child := range children[0] {
		for _, s := range childTraversed {
			traversed = append(traversed, string(child)+s)
		}
	}

	return traversed
}
