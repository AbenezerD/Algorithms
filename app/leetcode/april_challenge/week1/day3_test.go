package week1

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxSubArray(t *testing.T) {
	Convey("Given: ", t, func() {
		So(maxSubArray([]int{4,1,-2,3,5,7,-5,4}), ShouldEqual, 18)
		So(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}), ShouldEqual, 6)
		So(maxSubArray([]int{-2,1,-2,3,5,7,-5,6}), ShouldEqual, 16)
		So(maxSubArray([]int{-2}), ShouldEqual, -2)
	})
}
