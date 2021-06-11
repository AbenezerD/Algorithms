package medium

import "strconv"

/**
You are given two non-empty linked lists representing two non-negative integers.
The most significant digit comes first and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:
Input: l1 = [7,2,4,3], l2 = [5,6,4]
Output: [7,8,0,7]

Example 2:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [8,0,7]

Example 3:
Input: l1 = [0], l2 = [0]
Output: [0]

Constraints:
The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Revert := revertList(l2)
	l2Revert := revertList(l2)

	var sum *ListNode
	var node *ListNode
	carry := 0
	for ; l1Revert != nil && l2Revert != nil; {
		sumVal := carry + l1Revert.Val + l2Revert.Val
		node = &ListNode{
			Val: sumVal%10,
		}
		if sum == nil {
			sum = node
		} else {
			sum.Next = node
		}
		sum = node
		carry = sumVal/10
		l1Revert = l1Revert.Next
		l2Revert = l2Revert.Next
		sum = sum.Next
	}

	return revertList(sum)
}
// 1 2 3 4
// 9 2 7
// 0 5 0 5

// 4231
//  729
// 5050

func toString(l *ListNode) string {
	str := ""
	for ; l != nil; {
		str += strconv.Itoa(l.Val)
		l = l.Next
	}

	return str
}

func getLength(l *ListNode) int {
	count := 0
	for ; l != nil; {
		count++
		l = l.Next
	}

	return count
}

// 1 -> 2 -> 3 -> 4
// 1
// 2 -> 1
//
// 4 -> 3 -> 2 -> 1

func revertList(l *ListNode) *ListNode {
	var node *ListNode
	for ; l != nil; {
		temp := &ListNode{Val: l.Val}
		temp.Next = node
		node = temp
		l = l.Next
	}

	return node
}
