package medium

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseBetween(t *testing.T) {
	Convey("Test: reverseBetween", t, func() {
		So(reverseBetween(createNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), 2, 4), ShouldResemble,
			createNode([]int{1, 4, 3, 2, 5, 6, 7, 8, 9, 10}))
		So(reverseBetween(createNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), 8, 10), ShouldResemble,
			createNode([]int{1, 2, 3, 4, 5, 6, 7, 10, 9, 8}))
		So(reverseBetween(createNode([]int{1, 2, 3, 4, 5}), 2, 4), ShouldResemble, createNode([]int{1, 4, 3, 2, 5}))
		So(reverseBetween(createNode([]int{1, 2, 3, 4, 5}), 3, 4), ShouldResemble, createNode([]int{1, 2, 4, 3, 5}))
		So(reverseBetween(createNode([]int{1, 2, 3, 4, 5}), 2, 5), ShouldResemble, createNode([]int{1, 5, 4, 3, 2}))
		So(reverseBetween(createNode([]int{1, 2, 3, 4, 5}), 1, 5), ShouldResemble, createNode([]int{5, 4, 3, 2, 1}))
		So(reverseBetween(createNode([]int{1, 2, 3, 4, 5}), 1, 3), ShouldResemble, createNode([]int{3, 2, 1, 4, 5}))
		So(reverseBetween(createNode([]int{5}), 1, 1), ShouldResemble, createNode([]int{5}))
	})
}
