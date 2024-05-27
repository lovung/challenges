package may2024

import "github.com/lovung/ds/lists"

// https://leetcode.com/problems/double-a-number-represented-as-a-linked-list/
func doubleIt(head *lists.ListNode[int]) *lists.ListNode[int] {
	tail := reverseNodes(head)
	node, c := tail, 0
	for node != nil {
		val := node.Val*2 + c
		node.Val = val % 10
		c = val / 10
		node = node.Next
	}
	if c > 0 {
		newNode := lists.NewListNode(c)
		head.Next = &newNode
	}
	return reverseNodes(tail)
}
