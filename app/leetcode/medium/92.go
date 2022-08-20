package medium

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	node := &ListNode{Next: head}

	p1 := node
	p2 := node.Next
	for i, j := 1, 1; i < right; i++ {
		if j < left {
			j++
			p1 = p1.Next
			p2 = p2.Next
			continue
		}
		temp := p2.Next
		p2.Next = temp.Next
		temp.Next = p1.Next
		p1.Next = temp
		j++
	}

	return node.Next
}
