package medium

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTopKFrequent(t *testing.T) {
	Convey("Test: topKFrequent", t, func() {
		So(topKFrequent([]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}, 10),
			ShouldResemble, []int{1, 2, 5, 3, 6, 7, 4, 8, 10, 11})
		So(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2), ShouldResemble, []int{1, 2})
		So(topKFrequent([]int{1}, 1), ShouldResemble, []int{1})
	})
}
