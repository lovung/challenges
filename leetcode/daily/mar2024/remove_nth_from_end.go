package mar2024

import "github.com/lovung/ds/lists"

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *lists.ListNode[int], n int) *lists.ListNode[int] {
	s, f := (*lists.ListNode[int])(nil), head
	for range n {
		f = f.Next
	}
	for f != nil {
		if s == nil {
			s = head
		} else {
			s = s.Next
		}
		f = f.Next
	}
	if s == nil {
		return head.Next
	}
	if s.Next == nil {
		return nil
	}
	s.Next = s.Next.Next
	return head
}
