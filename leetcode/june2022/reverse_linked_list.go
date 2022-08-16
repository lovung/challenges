package june2022

// ListNode in linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Link: https://leetcode.com/problems/reverse-linked-list-ii/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
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
		prev, next       *ListNode
		preTail, subHead *ListNode
		newHead          *ListNode
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

func reverseList(head *ListNode) *ListNode {
	var (
		runner     = head
		prev, next *ListNode
	)

	for runner != nil {
		next = runner.Next
		runner.Next = prev
		prev, runner = runner, next
	}
	return prev
}

func swapNodes(head *ListNode, k int) *ListNode {
	runner := head
	n := 0
	for runner != nil {
		n++
		runner = runner.Next
	}
	runner = head
	i := 0
	var first *ListNode
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
