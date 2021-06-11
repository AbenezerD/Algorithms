package week2

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFindMaxLength(t *testing.T) {
	Convey("Test: LastStoneWeight ", t, func() {
		So(findMaxLength([]int{0,1}), ShouldEqual, 2)
		So(findMaxLength([]int{0,1,0}), ShouldEqual, 2)
		So(findMaxLength([]int{0,1,0,0}), ShouldEqual, 2)
		So(findMaxLength([]int{0,0,0,0}), ShouldEqual, 0)
		So(findMaxLength([]int{1,1,1,1}), ShouldEqual, 0)
		So(findMaxLength([]int{0,1,1,0,1,0,1}), ShouldEqual, 6)
		So(findMaxLength([]int{0,0,1,1,0,1,0,1}), ShouldEqual, 8)
		So(findMaxLength([]int{0,1,0,0,0,0,1,1,0,1,0,1,0,0,0,0,1}), ShouldEqual, 8)
		So(findMaxLength([]int{0,1,0,0,0,0,1,1,0,1,0,1,0,0,0,0,0,1,0,0,1,0,0,0,1,1,1,1}), ShouldEqual, 12)
	})
}
