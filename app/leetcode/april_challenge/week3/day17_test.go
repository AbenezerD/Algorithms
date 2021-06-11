package week3

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNumIslands(t *testing.T) {
	Convey("Test: NumIslands", t, func() {
		So(numIslands([][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}), ShouldEqual, 1)
		So(numIslands([][]byte{
			{1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 1},
		}), ShouldEqual, 3)
		So(numIslands([][]byte{
			{0, 0, 1, 0, 0},
			{1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}), ShouldEqual, 1)
	})
}
