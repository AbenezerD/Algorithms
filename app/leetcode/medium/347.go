package medium

import "sort"

func topKFrequent(nums []int, k int) []int {
	countIndex := make(map[int]int, 0)

	for _, num := range nums {
		countIndex[num] += 1
	}

	counts := make([]int, 0)
	reverseIndex := make(map[int][]int, 0)
	for num, count := range countIndex {
		if _, ok := reverseIndex[count]; !ok {
			counts = append(counts, count)
		}
		reverseIndex[count] = append(reverseIndex[count], num)
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})
	frequents := make([]int, 0)
	for _, count := range counts {
		if len(frequents) >= k {
			break
		}
		frequents = append(frequents, reverseIndex[count]...)
	}
	return frequents
}
