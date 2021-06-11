package week1

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCountElements(t *testing.T) {
	Convey("Given: ", t, func() {
		So(countElements([]int{1,2,3}), ShouldEqual, 2)
		So(countElements([]int{1,1,3,3,5,5,7,7}), ShouldEqual, 0)
		So(countElements([]int{1,3,2,3,5,0}), ShouldEqual, 3)
		So(countElements([]int{1,1,2,2}), ShouldEqual, 2)
	})
}
