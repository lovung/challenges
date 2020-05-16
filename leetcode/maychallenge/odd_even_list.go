package maychallenge

func oddEvenList(head *ListNode) *ListNode {
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
