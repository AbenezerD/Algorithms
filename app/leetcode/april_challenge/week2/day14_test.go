package week2

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStringShift(t *testing.T) {
	Convey("Test: StringShift", t, func() {
		So(stringShift("abc", [][]int{{0,1},{1,2}}), ShouldEqual, "cab")
		So(stringShift("abcdefg", [][]int{{1,1},{1,1},{0,2},{1,3}}), ShouldEqual, "efgabcd")
	})
}
