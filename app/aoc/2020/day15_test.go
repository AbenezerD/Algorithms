package main

import (
	"testing"
)

func TestDay15(t *testing.T) {
	Convey("Given: ", t, func() {
		So(day15([]int{0, 3, 6}), ShouldEqual, 436)
		So(day15([]int{1, 3, 2}), ShouldEqual, 1)
		So(day15([]int{1, 2, 3}), ShouldEqual, 27)
		So(day15([]int{2, 1, 3}), ShouldEqual, 10)
		So(day15([]int{2, 3, 1}), ShouldEqual, 78)
		So(day15([]int{3, 2, 1}), ShouldEqual, 438)
		So(day15([]int{3, 1, 2}), ShouldEqual, 1836)
		So(day15([]int{0, 1, 5, 10, 3, 12, 19}), ShouldEqual, 1373)
	})
}

func TestDay15b(t *testing.T) {
	Convey("Given: ", t, func() {
		So(day15b([]int{0, 3, 6}), ShouldEqual, 175594)
		So(day15b([]int{1, 3, 2}), ShouldEqual, 2578)
		So(day15b([]int{1, 2, 3}), ShouldEqual, 261214)
		So(day15b([]int{2, 1, 3}), ShouldEqual, 3544142)
		So(day15b([]int{2, 3, 1}), ShouldEqual, 6895259)
		So(day15b([]int{3, 2, 1}), ShouldEqual, 18)
		So(day15b([]int{3, 1, 2}), ShouldEqual, 362)
		So(day15b([]int{0, 1, 5, 10, 3, 12, 19}), ShouldEqual, 112458)
	})
}
