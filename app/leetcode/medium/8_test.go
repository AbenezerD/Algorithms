package medium

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMyAtoi(t *testing.T) {
	Convey("Test: myAtoi", t, func() {
		So(myAtoi("42"), ShouldEqual, 42)
		So(myAtoi("   -42"), ShouldEqual, -42)
		So(myAtoi("4193 with words"), ShouldEqual, 4193)
		So(myAtoi("   -42"), ShouldEqual, -42)
		So(myAtoi("words and 987"), ShouldEqual, 0)
		So(myAtoi("-91283472332"), ShouldEqual, -2147483648)
	})
}
