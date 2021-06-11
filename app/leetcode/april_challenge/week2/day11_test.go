package week2

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	Convey("Test DiameterOfBinaryTree", t, func() {
		So(diameterOfBinaryTree(createTree()), ShouldEqual, 3)
		So(diameterOfBinaryTree(createTree1()), ShouldEqual, 2)
		So(diameterOfBinaryTree(createTree2()), ShouldEqual, 1)
	})
}

func createTree() *TreeNode {
	// (1, (2, (4, 5), 3), 5)
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
}

func createTree1() *TreeNode {
	// (1, (2, (4, 5), 3), 5)
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
}
func createTree2() *TreeNode {
	// (1, (2, (4, 5), 3), 5)
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
}
