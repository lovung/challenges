package june2022

import (
	. "github.com/lovung/challenges/internal/linkedlist"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Link: https://leetcode.com/problems/reverse-linked-list-ii/
func reverseBetween(head *ListNode[int], left int, right int) *ListNode[int] {
	if head == nil {
		return nil
	}
	if left == right {
		return head
	}

	li, ri := left-1, right-1
	i := 0
	var (
		runner           = head
		prev, next       *ListNode[int]
		preTail, subHead *ListNode[int]
		newHead          *ListNode[int]
	)
	for runner != nil {
		next = runner.Next
		switch {
		case i < li:
			newHead = head
		case i == li:
			preTail = prev
			subHead = runner
			runner.Next = nil
		case i > li && i < ri:
			runner.Next = prev
		case i == ri:
			runner.Next = prev
			if preTail != nil {
				preTail.Next = runner
			} else {
				newHead = runner
			}
		case i == ri+1:
			subHead.Next = runner
		}
		prev, runner = runner, next
		i++
	}
	return newHead
}

// Link: https://leetcode.com/problems/reverse-linked-list/
func reverseList(head *ListNode[int]) *ListNode[int] {
	var (
		runner     = head
		prev, next *ListNode[int]
	)

	for runner != nil {
		next = runner.Next
		runner.Next = prev
		prev, runner = runner, next
	}
	return prev
}
