package medium

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxAreaOfIsland(t *testing.T) {
	Convey("Test: maxAreaOfIsland", t, func() {
		grid := [][]int{
			{0, 1, 0, 1},
			{0, 1, 0, 0},
			{0, 0, 0, 1},
			{0, 1, 1, 1},
		}
		So(maxAreaOfIsland(grid), ShouldEqual, 4)

		grid = [][]int{
			{0, 1, 0, 1},
			{0, 1, 1, 0},
			{1, 0, 0, 1},
			{1, 1, 1, 1},
		}
		So(maxAreaOfIsland(grid), ShouldEqual, 6)

		grid = [][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0},
		}
		So(maxAreaOfIsland(grid), ShouldEqual, 7)
		So(maxAreaOfIsland([][]int{{0, 0, 0, 0, 0, 0}}), ShouldEqual, 0)
		So(maxAreaOfIsland([][]int{{1, 1, 1, 1, 1, 1}}), ShouldEqual, 6)
		So(maxAreaOfIsland([][]int{{1, 1, 1, 1, 1, 1}}), ShouldEqual, 6)
	})
}

func TestBenchmark(t *testing.T) {
	res := testing.Benchmark(func(b *testing.B) {
		Convey("Test: benchmark", b, func() {
			grid := [][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0},
			}
			So(maxAreaOfIsland(grid), ShouldEqual, 7)
		})
	})

	fmt.Printf("%v", res)
}
