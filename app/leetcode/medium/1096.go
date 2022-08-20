package medium

import (
	"sort"
)

func braceExpansionII(expression string) []string {
	wordsIndex := make(map[string]struct{}, 0)

	var groups [][]string
	var group []string
	var current string
	openBraceCount := 0
	//var prev rune
	for _, c := range expression {
		switch c {
		case '{':
			if len(group) > 0 && openBraceCount <= 1 {
				groups = append(groups, group)
			}
			if current != "" && openBraceCount <= 1 {
				groups = append(groups, []string{current})
				current = ""
			}
			openBraceCount++
		case ',':
			if openBraceCount <= 1 {
				if len(groups) == 0 && current != "" {
					groups = append(groups, []string{current})
				}
				words := traverse96(groups)
				for _, word := range words {
					wordsIndex[word] = struct{}{}
				}
				groups = nil
				group = nil
				current = ""
			} else {
				group = append(group, current)
				current = ""
			}
		case '}':
			openBraceCount--
			if len(group) > 0 && openBraceCount <= 1 {
				group = append(group, current)
				groups = append(groups, group)
				current = ""
				group = nil
			}
		default:
			current += string(c)
		}
	}

	if len(group) > 0 {
		groups = append(groups, group)
	}
	if current != "" {
		groups = append(groups, []string{current})
	}

	if len(groups) > 0 {
		words := traverse96(groups)
		for _, word := range words {
			wordsIndex[word] = struct{}{}
		}
	}

	words := make([]string, 0, len(wordsIndex))
	for word, _ := range wordsIndex {
		words = append(words, word)
	}
	sort.Strings(words)
	return words
}

func traverse96(children [][]string) []string {
	if len(children) == 0 {
		return nil
	}
	if len(children) == 1 {
		strArr := make([]string, len(children[0]))
		for i, child := range children[0] {
			strArr[i] = child
		}
		return strArr
	}

	var traversed []string
	childTraversed := traverse96(children[1:])
	for _, child := range children[0] {
		for _, s := range childTraversed {
			traversed = append(traversed, child+s)
		}
	}

	return traversed
}
