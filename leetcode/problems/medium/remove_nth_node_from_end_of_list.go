package medium

import "github.com/lovung/ds/lists"

func removeNthFromEnd(head *lists.ListNode[int], n int) *lists.ListNode[int] {
	var second *lists.ListNode[int]
	count := 0
	for first := head; first != nil; first = first.Next {
		if count < n {
			count++
			continue
		}
		if second == nil {
			second = head
		} else {
			second = second.Next
		}
	}
	if second == nil {
		return head.Next
	}
	second.Next = second.Next.Next
	return head
}
