package medium

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestZigZagConversion(t *testing.T) {
	Convey("Test: zigZagConversion", t, func() {
		So(convert("A", 1), ShouldEqual, "A")
		So(convert("ABCD", 1), ShouldEqual, "ABCD")
		So(convert("ABCD", 2), ShouldEqual, "ACBD")
		So(convert("ABCD", 3), ShouldEqual, "ABDC")
		So(convert("ABCD", 4), ShouldEqual, "ABCD")
		So(convert("PAYPALISHIRING", 3), ShouldEqual, "PAHNAPLSIIGYIR")
		So(convert("PAYPALISHIRING", 4), ShouldEqual, "PINALSIGYAHRPI")
	})
}
