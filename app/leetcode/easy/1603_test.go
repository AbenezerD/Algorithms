package easy

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAverage(t *testing.T) {
	Convey("Test: average", t, func() {
		So(average([]int{1, 2, 3}), ShouldEqual, 2.0)
		So(average([]int{1, 2, 3, 4}), ShouldEqual, 2.5)
		So(average([]int{1000, 2000, 3000, 4000}), ShouldEqual, 2500)
		So(average([]int{1000000, 2000000, 3000000, 4000000}), ShouldEqual, 2500000)
	})
}
