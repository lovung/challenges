package microsoft

import "github.com/lovung/ds/lists"

// Given a linked list, remove the n-th node from the end of list and return its head.

// Type of Value in Node: int.
// Can the values duplicated: yes.
// 0 < n <= len(linked list)

//      1 -> 2 -> 3 -> 4 | n: 2
// l              r
//      l              r
//           l              r=nil

//      1 -> 2 -> 3 -> 4 | n: 3
// l                   r
//      l                   r=nil
//
//      1 -> 2 -> 3 -> 4 | n: 4
// l                        r=nil
//           2 -> 3 -> 4
//
// 1 -> 2 -> 4
// * 1 | n: 1 -> null
// * nil | n: 0 -> null

// Solution:
// r - l == n+1

func RemoveNThNode(head *lists.ListNode[int], n int) *lists.ListNode[int] {
	if head == nil {
		return nil
	}
	l, r := (*lists.ListNode[int])(nil), head
	for range n {
		r = r.Next
	}
	for r != nil {
		if l == nil {
			l = head
		} else {
			l = l.Next
		}
		r = r.Next
	}
	if l == nil {
		// Remove head node here
		newHead := head.Next
		head.Next = nil
		return newHead
	}
	// Remove l.Next node here
	l.Next = l.Next.Next
	return head
}
