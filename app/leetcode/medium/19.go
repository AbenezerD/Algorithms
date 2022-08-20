package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	index := make(map[int]*ListNode, 0)

	node := head
	i := 0
	for node != nil {
		index[i] = node
		node = node.Next
		i++
	}

	if i == n {
		head = head.Next
	} else {
		index[i-n-1].Next = index[i-n+1]
	}

	return head
}
