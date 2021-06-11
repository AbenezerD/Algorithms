package week3

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBstFromTraversal(t *testing.T) {
	Convey("Test: bstFromTraversal", t, func() {
		So(bstFromPreorder([]int{8, 5, 1, 7, 10, 9}), ShouldResemble, createTreeNode())
	})
}

func createTreeNode() *TreeNode {
	return &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 9,
			},
		},
	}
}
