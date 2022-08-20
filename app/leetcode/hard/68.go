package hard

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	var groups []string
	currentGroup := words[0]
	for i := 1; i < len(words); i++ {
		word := words[i]
		next := currentGroup + " " + word
		if len(next) <= maxWidth {
			currentGroup = next
			continue
		}

		group := strings.Split(currentGroup, " ")
		padding, extra := getPaddingAndExtras(currentGroup, maxWidth)
		line := group[0]
		if len(group) == 1 {
			line += padding
		} else {
			for i := 1; i < len(group); i++ {
				line += padding
				if i <= extra {
					line += " "
				}
				line += group[i]
			}
		}

		groups = append(groups, line)
		currentGroup = word
	}
	spaceLeft := maxWidth - len(currentGroup)
	groups = append(groups, currentGroup+getNSpaces(spaceLeft))

	return groups
}

func getPaddingAndExtras(currentGroup string, maxWidth int) (string, int) {
	group := strings.Split(currentGroup, " ")
	if len(group) == 1 {
		return getNSpaces(maxWidth - len(currentGroup)), 0
	}

	groupLen := len(group)
	spaceLeft := maxWidth - len(currentGroup) + groupLen - 1
	even := spaceLeft / (groupLen - 1)
	evenSpace := getNSpaces(even)
	pad := spaceLeft % (groupLen - 1)

	return evenSpace, pad
}

func getNSpaces(n int) string {
	space := ""
	for i := 0; i < n; i++ {
		space += " "
	}

	return space
}
