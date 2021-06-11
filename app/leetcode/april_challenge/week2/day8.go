package week2

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	occurrence := make(map[int]*ListNode, 0)
	node := head
	index := 0
	for ;node != nil; {
		occurrence[index] = node
		node = node.Next
		index++
	}

	return occurrence[index/2]
}
