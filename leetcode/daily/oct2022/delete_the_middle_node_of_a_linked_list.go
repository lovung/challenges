package oct2022

import "github.com/lovung/ds/lists"

// l: need remove the l.next
func deleteMiddle(head *lists.ListNode[int]) *lists.ListNode[int] {
	if head == nil || head.Next == nil {
		return nil
	}
	l, r := head, head
	if r.Next != nil {
		r = r.Next.Next
	}
	for r != nil && r.Next != nil {
		r = r.Next.Next
		l = l.Next
	}

	l.Next = l.Next.Next
	return head
}
