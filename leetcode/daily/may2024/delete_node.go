package may2024

import "github.com/lovung/ds/lists"

// https://leetcode.com/problems/delete-node-in-a-linked-list/description/
func deleteNode(node *lists.ListNode[int]) {
	node.Val, node.Next = node.Next.Val, node.Next.Next
}
