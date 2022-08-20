package medium

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubsets(t *testing.T) {
	Convey("Test: subsets", t, func() {
		So(subsets([]int{1, 2, 3}), ShouldResemble, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}})
	})
}
