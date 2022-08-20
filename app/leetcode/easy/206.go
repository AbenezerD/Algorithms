package easy

func reverseList(head *ListNode) *ListNode {
	node := &ListNode{Next: head}

	p1 := node
	p2 := node.Next
	for p2 != nil && p2.Next != nil {
		temp := p2.Next
		p2.Next = temp.Next
		temp.Next = p1.Next
		p1.Next = temp
	}

	return node.Next
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	h, _ := recursive(head)
	return h
}

func recursive(node *ListNode) (*ListNode, *ListNode) {
	if node.Next == nil {
		return node, node
	}
	n := &ListNode{Val: node.Val}
	tempHead, tempTail := recursive(node.Next)
	tempTail.Next = n
	return tempHead, n
}
