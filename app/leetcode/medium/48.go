package medium

func rotateInformal(matrix [][]int) [][]int {
	rotated := make([][]int, len(matrix))
	for i := range rotated {
		rotated[i] = make([]int, len(matrix[i]))
	}

	for i, arrays := range matrix {
		n := len(arrays)
		for j := 0; j < n; j++ {
			rotated[i][j] = matrix[n-1-j][i]
		}
	}

	return rotated
}

func rotate(matrix [][]int) [][]int {
	n := len(matrix)
	holdLen := (n*n + 1) / 2
	hold := make([]int, 0, holdLen)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if j >= i {
				hold = append(hold, matrix[i][j])
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j < n-1-i {
				matrix[i][j] = matrix[n-1-j][i]
			} else {
				k := (n - j - 1) * (n - j) / 2
				i2 := (n * (n - j - 1)) + i - k
				matrix[i][j] = hold[i2]
			}
		}
	}

	return matrix
}

// 1,2,3
// 4,5,6
// 7,8,9

// 1,4,7  ->  7,4,1
// 2,5,8  ->  8,5,2
// 3,6,9  ->  9,6,3

func rotate3(matrix [][]int) [][]int {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[i][n-j-1]
			matrix[i][n-j-1] = temp
		}
	}

	return matrix
}
