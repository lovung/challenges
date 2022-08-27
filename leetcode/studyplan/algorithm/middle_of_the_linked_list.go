package algorithm

import (
	"github.com/lovung/ds/lists"
)

// Link: https://leetcode.com/problems/middle-of-the-linked-list/
// BigO: O(n)
func middleNode(head *lists.ListNode[int]) *lists.ListNode[int] {
	first, second := head, head
	blackRed := false
	for first != nil {
		if blackRed {
			second = second.Next
		}
		blackRed = !blackRed
		first = first.Next
	}
	return second
}
