package medium

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindMedianSortedArrays(t *testing.T) {
	Convey("Test: findMedianSortedArrays", t, func() {
		So(findMedianSortedArrays([]int{0, 0}, []int{0, 0}), ShouldEqual, 0)
		So(findMedianSortedArrays([]int{1, 3, 9, 19}, []int{2, 7, 90}), ShouldEqual, 7)
		So(findMedianSortedArrays([]int{1, 3}, []int{2}), ShouldEqual, 2)
		So(findMedianSortedArrays([]int{1, 2}, []int{3, 4}), ShouldEqual, 2.5)
	})
}
