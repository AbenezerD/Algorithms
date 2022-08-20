package easy

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMerge(t *testing.T) {
	Convey("Test: merge", t, func() {
		//So(merge([]int{1, 3, 7, 0, 0, 0, 0}, 3, []int{2, 5, 7, 8}, 4), ShouldResemble, []int{1, 2, 3, 5, 7, 7, 8})
		//So(merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3), ShouldResemble, []int{1, 2, 2, 3, 5, 6})
		//So(merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{4, 5, 6}, 3), ShouldResemble, []int{1, 2, 3, 4, 5, 6})
		So(merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3), ShouldResemble, []int{1, 2, 3, 4, 5, 6})
		//So(merge([]int{1}, 1, []int{}, 0), ShouldResemble, []int{1})
		//So(merge([]int{0}, 0, []int{1}, 1), ShouldResemble, []int{1})
		//So(merge([]int{1, 0, 0}, 1, []int{2, 3}, 2), ShouldResemble, []int{1, 2, 3})
		//So(merge([]int{1, 0}, 1, []int{2}, 1), ShouldResemble, []int{1, 2})
	})
}

//[0]
//0
//[1]
//1
