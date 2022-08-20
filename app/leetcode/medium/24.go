package medium

func swapPairs(head *ListNode) *ListNode {
	//node := head
	node := head
	temp := head.Next
	node.Next = temp.Next
	temp.Next = node
	head = temp
	node = head.Next
	n := head.Next.Next
	for n != nil && n.Next != nil {
		temp := n.Next
		n.Next = temp.Next
		temp.Next = n
		node.Next = temp
		n = n.Next
	}

	return head
}
