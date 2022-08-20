package medium

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAllPathsSourceTarget(t *testing.T) {
	Convey("Test: allPathsSourceTarget", t, func() {
		So(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}), ShouldResemble, [][]int{{0, 1, 3}, {0, 2, 3}})
		So(allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}), ShouldResemble,
			[][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}})
		So(allPathsSourceTarget([][]int{{1, 2, 3}, {2}, {3}, {}}), ShouldResemble,
			[][]int{{0, 1, 2, 3}, {0, 2, 3}, {0, 3}})
		So(allPathsSourceTarget([][]int{{1, 3}, {2}, {3}, {}}), ShouldResemble,
			[][]int{{0, 1, 2, 3}, {0, 3}})
		So(allPathsSourceTarget([][]int{{2}, {}, {1}}), ShouldResemble,
			[][]int{{0, 2}})
	})
}
