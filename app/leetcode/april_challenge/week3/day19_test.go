package week3

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSearch(t *testing.T) {
	Convey("Test: Search", t, func() {
		So(search([]int{1}, 2), ShouldEqual, -1)
		So(search([]int{1}, 1), ShouldEqual, 0)
		So(search([]int{4, 5, 6, 7, 0, 1, 2}, 3), ShouldEqual, -1)
		So(search([]int{4, 5, 6, 7, 0, 1, 2}, 4), ShouldEqual, 0)
		So(search([]int{4, 5, 6, 7, 0, 1, 2}, 5), ShouldEqual, 1)
		So(search([]int{4, 5, 6, 7, 0, 1, 2}, 6), ShouldEqual, 2)
		So(search([]int{4, 5, 6, 7, 0, 1, 2}, 7), ShouldEqual, 3)
		So(search([]int{4, 5, 6, 7, 0, 1, 2}, 0), ShouldEqual, 4)
		So(search([]int{4, 5, 6, 7, 0, 1, 2}, 1), ShouldEqual, 5)
		So(search([]int{4, 5, 6, 7, 0, 1, 2}, 2), ShouldEqual, 6)
	})
}
