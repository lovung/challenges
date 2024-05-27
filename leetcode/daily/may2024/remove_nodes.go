package may2024

import "github.com/lovung/ds/lists"

// https://leetcode.com/problems/remove-nodes-from-linked-list
func removeNodes(head *lists.ListNode[int]) *lists.ListNode[int] {
	tail := reverseNodes(head)
	node, prev := tail, (*lists.ListNode[int])(nil)
	maxVal := 0
	for node != nil {
		if node.Val < maxVal {
			// Should be remove
			node, prev.Next = node.Next, node.Next
		} else {
			maxVal = max(maxVal, node.Val)
			node, prev = node.Next, node
		}
	}
	return reverseNodes(tail)
}

func reverseNodes(head *lists.ListNode[int]) *lists.ListNode[int] {
	node, prev := head, (*lists.ListNode[int])(nil)
	for node != nil {
		node, prev, node.Next = node.Next, node, prev
	}
	return prev
}
