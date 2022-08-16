package algorithm

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	first, second := head, head
	blackRed := false
	for first != nil {
		if blackRed {
			second = second.Next
		}
		blackRed = !blackRed
		first = first.Next
	}
	return second
}
