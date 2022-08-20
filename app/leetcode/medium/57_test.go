package medium

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSortColors(t *testing.T) {
	Convey("Test: sortColors", t, func() {
		var arr []string
		arr2 := make([]string, 0)

		So(append(arr, arr2...), ShouldBeNil)
		So(sortColors2([]int{1, 2, 2, 2, 2, 0, 0, 0, 1, 1}), ShouldResemble, []int{0, 0, 0, 1, 1, 1, 2, 2, 2, 2})
		So(sortColors2([]int{1, 2, 1, 0}), ShouldResemble, []int{0, 1, 1, 2})
		So(sortColors2([]int{1, 2, 1}), ShouldResemble, []int{1, 1, 2})
		So(sortColors2([]int{2, 1, 0}), ShouldResemble, []int{0, 1, 2})
		So(sortColors2([]int{2, 0, 1}), ShouldResemble, []int{0, 1, 2})
		So(sortColors2([]int{2, 0, 2, 1, 1, 0}), ShouldResemble, []int{0, 0, 1, 1, 2, 2})
		So(sortColors2([]int{0}), ShouldResemble, []int{0})
		So(sortColors2([]int{1}), ShouldResemble, []int{1})
		So(sortColors2([]int{0, 1, 0}), ShouldResemble, []int{0, 0, 1})
		So(sortColors2([]int{1, 2, 0}), ShouldResemble, []int{0, 1, 2})
		So(sortColors2([]int{0, 0, 0, 0, 0, 1, 2, 1, 2, 0, 0, 0, 0}), ShouldResemble,
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 2, 2})
	})
}
