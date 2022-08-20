package medium

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSwapPairs(t *testing.T) {
	Convey("Test: swapPairs", t, func() {
		So(swapPairs(createNode([]int{1, 2, 3, 4})), ShouldResemble, createNode([]int{2, 1, 4, 3}))
		So(swapPairs(createNode([]int{1, 2, 3, 4, 5})), ShouldResemble, createNode([]int{2, 1, 4, 3, 5}))
		So(swapPairs(createNode([]int{1, 2, 3, 4, 5, 6})), ShouldResemble, createNode([]int{2, 1, 4, 3, 6, 5}))
	})
}
