package problems

import (
	. "github.com/lovung/challenges/internal/linkedlist"
)

// Link: https://leetcode.com/problems/swapping-nodes-in-a-linked-list/
func swapNodes(head *ListNode[int], k int) *ListNode[int] {
	runner := head
	n := 0
	for runner != nil {
		n++
		runner = runner.Next
	}
	runner = head
	i := 0
	var first *ListNode[int]
	for runner != nil {
		if i == k-1 || i == n-k {
			if first == nil {
				first = runner
			} else {
				first.Val, runner.Val = runner.Val, first.Val
			}
		}
		i++
		runner = runner.Next
	}
	return head
}
