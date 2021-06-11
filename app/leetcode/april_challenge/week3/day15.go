package week3

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))

	product := 1
	zeroCount, zeroIndex := 0, 0
	for i, num := range nums {
		if num == 0 {
			zeroCount++
			zeroIndex = i
			if zeroCount > 1 {
				return output
			}
			continue
		}
		product *= num
	}

	if zeroCount == 1 {
		output[zeroIndex] = product
		return output
	}
	for i, num := range nums {
		output[i] = divide(product, num)
	}
	return output
}

func divide(dividend, divisor int) int {
	negative := dividend*divisor < 0
	if dividend < 0 {
		dividend *= -1
	}
	if divisor < 0 {
		divisor *= -1
	}
	low := 1
	high := dividend
	res := 0

	for ; low <= high; {
		mid := low + (high-low)/2
		x := divisor * mid
		if x <= dividend {
			res = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if negative {
		res *= -1
	}
	return res
}
