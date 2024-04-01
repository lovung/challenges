package mar2024

import "github.com/lovung/ds/lists"

// https://leetcode.com/problems/reorder-list
func reorderList(head *lists.ListNode[int]) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	tail := reverseList(slow)
	for l, r := head, tail; l != nil && r != nil && l.Next != r; {
		l.Next, r.Next, l, r = r, l.Next, l.Next, r.Next
	}
}

func reverseList(head *lists.ListNode[int]) *lists.ListNode[int] {
	cur, prev := head, (*lists.ListNode[int])(nil)

	for cur != nil {
		cur, cur.Next, prev = cur.Next, prev, cur
	}
	return prev
}
