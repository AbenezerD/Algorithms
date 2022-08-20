package medium

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAngleClock(t *testing.T) {
	Convey("Test: angleClock", t, func() {
		So(angleClock(12, 30), ShouldEqual, 165)
		So(angleClock(3, 30), ShouldEqual, 75)
		So(angleClock(3, 15), ShouldEqual, 7.5)
		So(angleClock(4, 50), ShouldEqual, 155)
		So(angleClock(12, 0), ShouldEqual, 0)
	})
}
