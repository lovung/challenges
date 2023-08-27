package contest

import (
	"github.com/lovung/ds/lists"
	"github.com/lovung/ds/stack"
)

// https://leetcode.com/problems/double-a-number-represented-as-a-linked-list/
func doubleIt(head *lists.ListNode[int]) *lists.ListNode[int] {
	s := stack.NewArrayStack[*lists.ListNode[int]](1e5)
	r := head
	for r != nil {
		s.Push(r)
		r = r.Next
	}
	c := 0
	for s.Len() > 0 {
		r = s.Pop()
		r.Val = r.Val*2 + c
		if r.Val >= 10 {
			r.Val -= 10
			c = 1
		} else {
			c = 0
		}
	}
	if c > 0 {
		newHead := &lists.ListNode[int]{
			Val:  1,
			Next: head,
		}
		return newHead
	}
	return head
}
