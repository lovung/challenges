package oct2022

import "github.com/lovung/ds/lists"

// Link: https://leetcode.com/problems/delete-node-in-a-linked-list/
func deleteNode(node *lists.ListNode[int]) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
