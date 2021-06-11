package week3

import "math"

// This is the BinaryMatrix's API interface.
// You should not implement it, or speculate about its implementation
type BinaryMatrix struct {
	matrix [][]int
}

func (this BinaryMatrix) Get(i, j int) int {
	return this.matrix[i][j]
}

func (this BinaryMatrix) Dimensions() []int {
	return []int{len(this.matrix), len(this.matrix[0])}
}

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dimensions := binaryMatrix.Dimensions()
	m, n := dimensions[0], dimensions[1]

	lowest := math.MaxInt8
	for i := 0; i < m; i++ {
		l := binarySearch(binaryMatrix, i, 0, n-1)
		if l != -1 && l < lowest {
			lowest = l
		}
	}
	if lowest == math.MaxInt8 {
		return -1
	}
	return lowest
}

func binarySearch(matrix BinaryMatrix, row, left, right int) int {
	if left == right {
		if matrix.Get(row, left) == 0 {
			return -1
		}
		return left
	}
	mid := (left + right) / 2
	if matrix.Get(row, mid) == 0 {
		return binarySearch(matrix, row, mid+1, right)
	}
	return binarySearch(matrix, row, left, mid)
}
