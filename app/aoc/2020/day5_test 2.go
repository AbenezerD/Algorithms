package main

import (
	"testing"
)

func TestDay5(t *testing.T) {
	Convey("Test: day5", t, func() {
		So(binary(0, 127, "FFFFFFF"), ShouldEqual, 0)
		So(binary(0, 127, "BBBBBBB"), ShouldEqual, 127)
		So(binary(0, 127, "BFFFBBF"), ShouldEqual, 70)
		So(binary(0, 127, "FFFBBBF"), ShouldEqual, 14)
		So(binary(0, 127, "BBFFBBF"), ShouldEqual, 102)
		So(binary(0, 7, "RRR"), ShouldEqual, 7)
		So(binary(0, 7, "RLR"), ShouldEqual, 5)
		So(binary(0, 7, "LLR"), ShouldEqual, 1)
		So(binary(0, 7, "RLL"), ShouldEqual, 4)
		So(binary(0, 7, "LRL"), ShouldEqual, 4)
	})
}
