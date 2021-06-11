package week3

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLeftMostColumnWithOne(t *testing.T) {
	Convey("Test: LeftMostColumnWithOne", t, func() {
		matrix := BinaryMatrix{
			matrix: [][]int{
				{0, 0, 0, 0, 0, 1, 1, 1, 1, 1}, // 5
				{0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, // 2
				{0, 0, 0, 0, 0, 0, 1, 1, 1, 1}, // 6
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, // 9
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // -1
			},
		}

		So(leftMostColumnWithOne(matrix), ShouldEqual, 2)
	})
}

func TestBinarySearch(t *testing.T) {
	Convey("Test: BinarySearch", t, func() {
		matrix := BinaryMatrix{
			matrix: [][]int{
				{0, 0, 0, 0, 0, 1, 1, 1, 1, 1}, // 5
				{0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, // 2
				{0, 0, 0, 0, 0, 0, 1, 1, 1, 1}, // 6
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, // 9
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // -1
			},
		}

		len := len(matrix.matrix[0]) - 1
		So(binarySearch(matrix, 0, 0, len), ShouldEqual, 5)
		So(binarySearch(matrix, 1, 0, len), ShouldEqual, 2)
		So(binarySearch(matrix, 2, 0, len), ShouldEqual, 6)
		So(binarySearch(matrix, 3, 0, len), ShouldEqual, 9)
		So(binarySearch(matrix, 4, 0, len), ShouldEqual, -1)
	})
}
