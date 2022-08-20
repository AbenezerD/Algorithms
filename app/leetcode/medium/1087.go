package medium

import "sort"

func expand(s string) []string {
	var groups [][]string
	var group []string
	var current string
	for _, c := range s {
		switch c {
		case '{':
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
			group = append(group, current)
			groups = append(groups, group)
			current = ""
			group = nil
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

	if len(groups) == 0 {
		return nil
	}

	traversed := traverse(groups)
	sort.Strings(traversed)
	return traversed
}

func traverse(children [][]string) []string {
	if len(children) == 1 {
		strArr := make([]string, len(children[0]))
		for i, char := range children[0] {
			strArr[i] = string(char)
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

func expand2(s string) []string {
	var groups [][]rune
	var group []rune
	var current rune
	for _, c := range s {
		switch c {
		case '{':
			if len(group) > 0 {
				groups = append(groups, group)
			}
			if current > 0 {
				groups = append(groups, []rune{current})
				current = 0
			}
		case ',':
			group = append(group, current)
			current = 0
		case '}':
			group = append(group, current)
			groups = append(groups, group)
			current = 0
			group = nil
		default:
			current = c
		}
	}

	if len(group) > 0 {
		groups = append(groups, group)
	}
	if current > 0 {
		groups = append(groups, []rune{current})
	}

	if len(groups) == 0 {
		return nil
	}

	return traverse2(groups)
}

func traverse2(children [][]rune) []string {
	if len(children) == 1 {
		strArr := make([]string, len(children[0]))
		for i, char := range children[0] {
			strArr[i] = string(char)
		}
		return strArr
	}

	var traversed []string
	childTraversed := traverse2(children[1:])
	for _, child := range children[0] {
		for _, s := range childTraversed {
			traversed = append(traversed, string(child)+s)
		}
	}

	return traversed
}
