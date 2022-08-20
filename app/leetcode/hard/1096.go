package hard

import (
	"sort"
)

func splitExpressions(expression string) []string {
	var expressions []string
	braceCount := 0
	lastIndex := 0
	//var prev string
	//var prev2 string
	current := ""
	for i, c := range expression {
		switch c {
		case '{':
			{
				braceCount++
				current = "{"
			}
		case '}':
			{
				braceCount--
				current = "}"
			}
		case ',':
			{
				if braceCount == 1 {
					expressions = append(expressions, expression[lastIndex:i])
					lastIndex = i + 1
				}
				current = "}"
			}
		default:
			current += string(c)
		}
		//prev2 = prev
		//prev = current
	}
	if lastIndex < len(expression) {
		expressions = append(expressions, expression[lastIndex:])
	}

	return expressions
}

func braceExpansionII(expression string) []string {
	//if strings.HasPrefix(expression, "{{") {
	//	expression = strings.TrimPrefix(expression, "{")
	//	expression = strings.TrimSuffix(expression, "}")
	//}

	var expressions []string
	braceCount := 0
	lastIndex := 0
	//var prev rune
	//var prev2 rune
	current := ""
	for i, c := range expression {
		switch c {
		case '{':
			{
				braceCount++
				current = ""
			}
		case '}':
			{
				braceCount--
				current = ""
			}
		case ',':
			{
				if braceCount == 1 {
					expressions = append(expressions, expression[lastIndex:i])
					lastIndex = i + 1
				}
			}
		default:
			current += string(c)

		}
		//prev2 = prev
		//prev = c
	}
	if lastIndex < len(expression) {
		expressions = append(expressions, expression[lastIndex:])
	}

	wordsIndex := make(map[string]struct{}, 0)
	for i := 0; i < len(expressions); i++ {
		e := expressions[i]
		if i != 0 {
			e = "{" + e
		}
		if i != len(expressions)-1 {
			e += "}"
		}
		words := expand2(e)
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

func _braceExpansionII(expression string) []string {
	//if strings.HasPrefix(expression, "{") {
	//	expression = strings.TrimPrefix(expression, "{")
	//	expression = strings.TrimSuffix(expression, "}")
	//}
	wordsIndex := make(map[string]struct{}, 0)

	var groups [][]string
	var group []string
	var current string
	openBraceCount := 0
	var prev rune
	for _, c := range expression {
		switch c {
		case '{':
			if len(group) > 0 && prev != ',' {
				groups = append(groups, group)
			} else if current != "" && prev != ',' {
				groups = append(groups, []string{current})
				current = ""
			}
			openBraceCount++
		case ',':
			if openBraceCount == 1 {
				if current != "" {
					group = append(group, current)
					groups = append(groups, group)
				}
				words := traverse(groups)
				for _, word := range words {
					wordsIndex[word] = struct{}{}
				}
				groups = nil
				group = nil
				current = ""
			} else {
				if current != "" {
					group = append(group, current)
					current = ""
				}
			}
		case '}':
			openBraceCount--
			if len(group) > 0 {
				group = append(group, current)
				groups = append(groups, group)
				current = ""
				group = nil
			}
		default:
			if prev == '}' {
				if len(group) > 0 && openBraceCount == 0 {
					group = append(group, current)
					groups = append(groups, group)
					current = ""
					group = nil
				}
			}
			current += string(c)
		}
		prev = c
	}

	if len(group) > 0 {
		groups = append(groups, group)
	}
	if current != "" {
		groups = append(groups, []string{current})
	}

	if len(groups) > 0 {
		words := traverse(groups)
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

func expand2(s string) []string {
	var groups [][]string
	var group []string
	var current string
	var prev rune
	for _, c := range s {
		switch c {
		case '{':
			if prev == ',' {
				break
			}
			if len(group) > 0 {
				groups = append(groups, group)
			}
			if current != "" {
				groups = append(groups, []string{current})
				current = ""
			}
		case ',':
			group = append(group, current)
			current = ""
		case '}':
			if len(group) > 0 {
				group = append(group, current)
				groups = append(groups, group)
				current = ""
				group = nil
			}
		default:
			current += string(c)
		}
		prev = c
	}

	if len(group) > 0 {
		groups = append(groups, group)
	}
	if current != "" {
		groups = append(groups, []string{current})
	}

	if len(groups) == 0 {
		return nil
	}

	traversed := traverse(groups)
	sort.Strings(traversed)
	return traversed
}

func traverse(children [][]string) []string {
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
	childTraversed := traverse(children[1:])
	for _, child := range children[0] {
		for _, s := range childTraversed {
			traversed = append(traversed, child+s)
		}
	}

	return traversed
}
