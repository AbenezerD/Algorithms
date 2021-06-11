package week4

import (
	"math"
	"math/bits"
)

func rangeBitwiseAnd(m int, n int) int {
	diff := n - m
	if bits.Len(uint(m)) != bits.Len(uint(n)) {
		return 0
	}
	b := bits.Len(uint(m))
	diffBits := bits.Len(uint(diff))
	shift := math.Pow(2.0, float64(b)) - math.Pow(2.0, float64(diffBits))
	return m & int(shift) & n
}

func rangeBitwiseAndBruteForce(m int, n int) int {
	if bits.Len(uint(m)) != bits.Len(uint(n)) {
		return 0
	}
	b := m
	for i := m + 1; i <= n; i++ {
		b &= i
	}
	return b
}

// 9
// 110
// 111
// 110

// 1110 8 + 2   ---
// 1011 8 + 2 + 1  ---
// 1100 8 + 4 --------
// 1000 8 ------------

// diff = 1

//21 10101 & 10110 -> 10000 +
//24 11000
// 21 + 32
//  110101
//
