package week2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	//visited := make(map[*TreeNode]struct{})
	//longest := 0
	//node := root
	//for ; len(visited) ; {
	//
	//}
	if root == nil || root.Left == nil && root.Right == nil {
		return 0
	}
	return height(root.Left, root.Right)
}

func height(left, right *TreeNode) int {
	count := 0
	if left != nil {
		count += height(left.Left, left.Right)
	}
	if right != nil {
		count += height(right.Left, right.Right)
	}
	if left == nil && right == nil {
		count += 1
	}

	return count

	//}
	//if right != nil {
	//	count += height(right.Left, right.Right)
	//} else {
	//	count += 1
	//}
	//return count
}
