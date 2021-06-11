package week1

import "fmt"

func isHappy(n int) bool {
	existingSums := make(map[int]struct{}, 0)
	for ;true; {
		n = getSumOfSquares(n)
		if _, ok := existingSums[n]; ok {
			return false
		}
		if n == 1 {
			return true
		}

		existingSums[n] = struct{}{}
	}
	return false
}

func getSumOfSquares(n int) int {
	sum := 0
	for q := n; q > 0; {
		r := q % 10
		q /= 10
		sum += r*r
	}
	fmt.Printf("%d ->", sum)
	return sum
}
