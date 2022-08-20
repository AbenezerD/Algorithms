package easy

import (
	"math"
)

func countPrimes2(n int) int {
	primeCount := 0
	for i := 1; i < n; i++ {
		if isPrime(i) {
			primeCount++
		}
	}
	return primeCount
}

func isPrime(n int) bool {
	switch n {
	case 1:
		return false
	case 2, 3:
		return true

	default:
		if n%2 == 0 {
			return false
		}

		sqrt := math.Sqrt(float64(n))
		for i := 3; i <= int(sqrt); i += 2 {
			if n%i == 0 {
				return false
			}
		}

		return true
	}
}

func countPrimes(n int) int {
	primes := make([]int, n)
	for i := 0; i < n; i++ {
		primes[i] = i
	}

	sqrt := math.Sqrt(float64(n))
	for i := 2; i <= int(sqrt); i++ {
		if isPrime(i) {
			for j := i * i; j < n; j += i {
				if primes[j]%i == 0 {
					primes[j] = -1
				}
			}
		}
	}

	count := 0
	for i := 2; i < len(primes); i++ {
		if primes[i] != -1 {
			count++
		}
	}
	return count
}
