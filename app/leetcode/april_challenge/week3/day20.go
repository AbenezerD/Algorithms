package week3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	var root *TreeNode
	for _, val := range preorder {
		root = insert(root, val)
	}
	return root
}

func insert(node *TreeNode, v int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: v}
	} else if v <= node.Val {
		node.Left = insert(node.Left, v)
	} else {
		node.Right = insert(node.Right, v)
	}
	return node
}
