package week4

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRangeBitwiseAnd(t *testing.T) {
	Convey("Test: RangeBitwiseAnd", t, func() {
		So(rangeBitwiseAnd(5, 6), ShouldEqual, rangeBitwiseAndBruteForce(5, 6))
		So(rangeBitwiseAnd(6, 7), ShouldEqual, rangeBitwiseAndBruteForce(6, 7))
		So(rangeBitwiseAnd(53, 56), ShouldEqual, rangeBitwiseAndBruteForce(53, 56))
		So(rangeBitwiseAnd(22, 24), ShouldEqual, rangeBitwiseAndBruteForce(22, 24))
		So(rangeBitwiseAnd(21, 24), ShouldEqual, rangeBitwiseAndBruteForce(21, 24))
		So(rangeBitwiseAnd(5, 7), ShouldEqual, 4)
		So(rangeBitwiseAnd(5, 15), ShouldEqual, 0)
		So(rangeBitwiseAnd(15, 19), ShouldEqual, 0)
		So(rangeBitwiseAnd(8, 15), ShouldEqual, 8)
		So(rangeBitwiseAnd(104, 126), ShouldEqual, 96)
		So(rangeBitwiseAnd(104, 127), ShouldEqual, 96)
		So(rangeBitwiseAnd(104, 106), ShouldEqual, rangeBitwiseAndBruteForce(104, 106))
		So(rangeBitwiseAnd(123, 127), ShouldEqual, rangeBitwiseAndBruteForce(123, 127))
		So(rangeBitwiseAnd(3, 5), ShouldEqual, 0)
	})
}
