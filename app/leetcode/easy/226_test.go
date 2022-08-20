package easy

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseBetween(t *testing.T) {
	Convey("Test: reverseList", t, func() {
		So(reverseListRecursive(createNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})), ShouldResemble,
			createNode([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
		So(reverseListRecursive(createNode([]int{1, 2, 3, 4, 5})), ShouldResemble, createNode([]int{5, 4, 3, 2, 1}))
		So(reverseListRecursive(createNode([]int{1, 2, 3, 4})), ShouldResemble, createNode([]int{4, 3, 2, 1}))
		So(reverseListRecursive(createNode([]int{1, 2, 3})), ShouldResemble, createNode([]int{3, 2, 1}))
		So(reverseListRecursive(createNode([]int{5})), ShouldResemble, createNode([]int{5}))
		So(reverseListRecursive(createNode([]int{})), ShouldResemble, createNode([]int{}))
	})
}
