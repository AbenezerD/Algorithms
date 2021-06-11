package week1

import "golang.org/x/tools/container/intsets"

func maxProfit(prices []int) int {
	localMax := 0
	localMin := intsets.MaxInt

	profit := 0
	for _, price := range prices {
		if localMax > localMin && price < localMax {
			profit += localMax - localMin
			localMin = intsets.MaxInt
			localMax = 0
		}
		if price < localMin {
			localMin = price
		}
		if price > localMax && price > localMin {
			localMax = price
		}
	}

	if localMax > localMin {
		profit += localMax - localMin
	}

	return profit
}
