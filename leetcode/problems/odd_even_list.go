package problems

import (
	. "github.com/lovung/ds/lists"
)

/*
 * Link: https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3331/
 */

func oddEvenList(head *ListNode[int]) *ListNode[int] {
	if head == nil {
		return head
	}
	head2 := head.Next
	if head2 == nil {
		return head
	}
	odd, even := head, head2
	for odd != nil && even != nil {
		odd.Next = even.Next
		if odd.Next == nil {
			break
		}
		even.Next = odd.Next.Next
		odd = odd.Next
		even = even.Next
	}
	odd.Next = head2
	return head
}
