package medium

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddTwoNumbers(t *testing.T) {
	Convey("Test: addTwoNumbers", t, func() {
		//Input: l1 = [7,2,4,3], l2 = [5,6,4]
		//Output: [7,8,0,7]

		l1 := createNode([]int{7, 2, 4, 3})
		l2 := createNode([]int{5, 6, 4})
		expected := createNode([]int{7, 8, 0, 7})

		sum := addTwoNumbers(l1, l2)
		So(sum, ShouldResemble, expected)
		So(toString(expected), ShouldEqual, "7807")

		So(addTwoNumbers(createNode([]int{5}), createNode([]int{5})), ShouldResemble, createNode([]int{1, 0}))
		So(addTwoNumbers(createNode([]int{1}), createNode([]int{9, 9})), ShouldResemble, createNode([]int{1, 0, 0}))
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
