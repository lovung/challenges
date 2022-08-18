package algorithm

import (
	. "github.com/lovung/challenges/internal/linkedlist"
)

// Link: https://leetcode.com/problems/middle-of-the-linked-list/
// BigO: O(n)
func middleNode(head *ListNode[int]) *ListNode[int] {
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
