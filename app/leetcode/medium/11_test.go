package medium

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxArea(t *testing.T) {
	Convey("Test: maxArea", t, func() {
		So(maxArea([]int{1, 8, 8, 2, 5, 4, 6, 5, 7, 8}), ShouldEqual, 64)
		So(maxArea([]int{1, 1}), ShouldEqual, 1)
		So(maxArea([]int{1, 2, 1}), ShouldEqual, 2)
		So(maxArea([]int{46, 53, 93, 82, 78, 20, 49, 76, 43, 67, 52, 41, 65, 38, 97, 98}), ShouldEqual, 1209)
	})
}
