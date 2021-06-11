package week3

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestProductExceptItself(t *testing.T) {
	Convey("Test: ProductExceptItself", t, func() {
		So(productExceptSelf([]int{1,2,3,4}), ShouldResemble, []int{24,12,8,6})
		So(productExceptSelf([]int{1,0}), ShouldResemble, []int{0,1})
		So(productExceptSelf([]int{0,1,0}), ShouldResemble, []int{0,0,0})
		So(divide(10, 3), ShouldEqual, 3)
		So(divide(26, 5), ShouldEqual, 5)
		So(divide(-10, 3), ShouldEqual, -3)
		So(divide(26, -5), ShouldEqual, -5)
		So(divide(-26, -5), ShouldEqual, 5)
	})
}
