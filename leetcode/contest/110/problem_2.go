package contest

import (
	"github.com/lovung/ds/lists"
	"github.com/lovung/ds/maths"
)

// https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list/
func insertGreatestCommonDivisors(head *lists.ListNode[int]) *lists.ListNode[int] {
	r := head
	for r != nil {
		next := r.Next
		if next == nil {
			break
		}
		newNode := lists.ListNode[int]{
			Val:  maths.GCD(r.Val, next.Val),
			Next: next,
		}
		r.Next = &newNode
		r = next
	}
	return head
}
