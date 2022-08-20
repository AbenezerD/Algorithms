package easy

import "fmt"

func readBinaryWatch(turnedOn int) []string {
	index := make(map[int][]int, 0)
	for i := 0; i < 60; i++ {
		indexCount := countOnes(i)
		index[indexCount] = append(index[indexCount], i)
	}

	maxHours := 3
	if turnedOn < 3 {
		maxHours = turnedOn
	}

	times := make([]string, 0)
	for i := 0; i <= maxHours; i++ {
		hours := index[i]
		for _, hour := range hours {
			if hour < 12 {
				minutes := index[turnedOn-i]
				for _, minute := range minutes {
					times = append(times, fmt.Sprintf("%d:%02d", hour, minute))
				}
			}
		}
	}
	return times
}

func countOnes(n int) int {
	count := 0
	for n > 0 {
		n = n & (n - 1)
		count++
	}
	return count
}

func hammingWeight(n uint) int {
	count := 0
	for n > 0 {
		n = n & (n - 1)
		count++
	}
	return count
}
