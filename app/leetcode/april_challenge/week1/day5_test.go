package week1

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxProfit(t *testing.T) {
	Convey("Given: ", t, func() {
		So(maxProfit([]int{7,1,5,3,6,4}), ShouldEqual, 7)
		So(maxProfit([]int{1,2,3,4,5}), ShouldEqual, 4)
		So(maxProfit([]int{7,6,4,3,1}), ShouldEqual, 0)
	})
}
