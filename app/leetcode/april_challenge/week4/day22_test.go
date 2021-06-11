package week4

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubarraySum(t *testing.T) {
	Convey("Test: SubarraySum", t, func() {
		So(subarraySum([]int{1, 1, 1}, 2), ShouldEqual, 2)
		So(subarraySum([]int{1, 2, 1, 0, 3, 0}, 3), ShouldEqual, 7)
		So(subarraySum([]int{2, 2, -2, 0, 4, 0}, 2), ShouldEqual, 6)
	})
}
