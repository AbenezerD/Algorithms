package week2

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCountElements(t *testing.T) {
	Convey("Given: ", t, func() {
		So(middleNode(
			createHeadNode([]int{1, 2, 3, 4, 5})), ShouldResemble,
			createHeadNode([]int{3,4,5}),
		)
		So(middleNode(
			createHeadNode([]int{1, 2, 3, 4, 5, 6})), ShouldResemble,
			createHeadNode([]int{4,5,6}),
		)
	})
}

func createHeadNode(arr []int) *ListNode {
	head := &ListNode{
		Val: arr[0],
	}

	currentNode := head
	for i := 1; i < len(arr); i++ {
		currentNode.Next = &ListNode{
			Val: arr[i],
		}

		currentNode = currentNode.Next
	}

	n := head
	for ; n != nil; {
		fmt.Printf("%d -> ", n.Val)
		n = n.Next
	}
	fmt.Printf("\n")

	return head
}
