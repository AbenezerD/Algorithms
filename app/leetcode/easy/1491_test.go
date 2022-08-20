package easy

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeTwoLists(t *testing.T) {
	Convey("Test: mergeTwoLists", t, func() {
		So(mergeTwoLists(createNode([]int{1, 2, 3}), createNode([]int{1, 3, 5})), ShouldResemble, createNode([]int{1, 1, 2, 3, 3, 5}))
		So(mergeTwoLists(createNode([]int{}), createNode([]int{0})), ShouldResemble, createNode([]int{0}))
		So(mergeTwoLists(createNode([]int{}), createNode([]int{})), ShouldResemble, createNode([]int{}))
	})
}

func createNode(intArr []int) *ListNode {
	if len(intArr) == 0 {
		return nil
	}
	l := &ListNode{Val: intArr[0]}
	node := l
	for i := 1; i < len(intArr); i++ {
		node.Next = &ListNode{Val: intArr[i]}
		node = node.Next
	}

	return l
}
