package jul2024

import "github.com/lovung/ds/lists"

// https://leetcode.com/problems/merge-nodes-in-between-zeros/
func mergeNodes(head *lists.ListNode[int]) *lists.ListNode[int] {
	lastZero, node := head, head.Next
	for node != nil {
		if node.Val == 0 {
			if node.Next == nil {
				lastZero.Next = nil
				break
			}
			lastZero, lastZero.Next = node, node
		} else {
			lastZero.Val += node.Val
		}
		node = node.Next
	}
	return head
}
