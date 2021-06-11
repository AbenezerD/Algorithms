package main

import (
	"testing"
)

func TestDay1(t *testing.T) {
	Convey("Test: day1", t, func() {
		So(get2([]int{1721, 979, 366, 299, 675, 1456}), ShouldEqual, 514579)
		So(get3([]int{1721, 979, 366, 299, 675, 1456}), ShouldEqual, 241861950)

		n := 720149 + 731033 + 1466212 + 1388501 + 1349179 + 816228 + 923226 + 973294 + 1004417 + 1122441 + 948607 + 979730 + 1398951 + 1038383 + 1079070 + 1116758 + 1215939
		So(n, ShouldEqual, 18272118)
	})
}
