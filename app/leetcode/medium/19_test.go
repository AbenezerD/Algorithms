package medium

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveNthFromEnd(t *testing.T) {
	Convey("Test: removeNthFromEnd", t, func() {
		So(removeNthFromEnd(createNode([]int{1, 2, 3, 4, 5}), 2), ShouldResemble, createNode([]int{1, 2, 3, 5}))
		So(removeNthFromEnd(createNode([]int{1, 2}), 1), ShouldResemble, createNode([]int{1}))
		So(removeNthFromEnd(createNode([]int{1, 2}), 2), ShouldResemble, createNode([]int{2}))
		So(removeNthFromEnd(createNode([]int{1}), 1), ShouldResemble, createNode([]int{}))
	})
}
