package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRomanToInt(t *testing.T) {
	Convey("Test RomanToInt", t, func() {
		//So(romanToInt("III"), ShouldEqual, 3)
		//So(romanToInt("LVIII"), ShouldEqual, 58)
		So(romanToInt("MCMXCIV"), ShouldEqual, 1994)
	})
}
