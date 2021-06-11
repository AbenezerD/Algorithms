package week1

import "sort"

func groupAnagrams(strs []string) [][]string {
	occurrence := make(map[string][]string, 0)
	for _, str := range strs {
		sortedString := sortString(str)
		occurrence[sortedString] = append(occurrence[sortedString], str)
	}

	groupedAnagrams := make([][]string, 0)
	for _, value := range occurrence {
		groupedAnagrams = append(groupedAnagrams, value)
	}

	return groupedAnagrams
}

func sortString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}
