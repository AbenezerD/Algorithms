package medium

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRob(t *testing.T) {
	Convey("Test: rob", t, func() {
		So(robII([]int{0}), ShouldEqual, 0)
		So(robII([]int{0, 0}), ShouldEqual, 0)
		So(robII([]int{2, 3, 2}), ShouldEqual, 3)
		So(robII([]int{1, 3, 1, 3, 100}), ShouldEqual, 103)
		So(robII([]int{1, 3, 1, 5, 6, 3, 100}), ShouldEqual, 109) // 107, 11
		So(robII([]int{4, 5, 6, 4, 5, 12, 5}), ShouldEqual, 22)   // 107, 11
	})
}
