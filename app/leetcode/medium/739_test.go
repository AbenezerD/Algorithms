package medium

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDailyTemperatures(t *testing.T) {
	Convey("Test: dailyTemperatures", t, func() {
		So(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}), ShouldResemble, []int{1, 1, 4, 2, 1, 1, 0, 0})
		So(dailyTemperatures([]int{30, 40, 50, 60}), ShouldResemble, []int{1, 1, 1, 0})
		So(dailyTemperatures([]int{30, 60, 90}), ShouldResemble, []int{1, 1, 0})
	})
}
