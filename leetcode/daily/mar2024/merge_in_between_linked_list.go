package mar2024

import "github.com/lovung/ds/lists"

// https://leetcode.com/problems/merge-in-between-linked-lists
func mergeInBetween(
	list1 *lists.ListNode[int],
	a int, b int,
	list2 *lists.ListNode[int],
) *lists.ListNode[int] {
	var endList2 *lists.ListNode[int]
	for n := list2; n != nil; n = n.Next {
		endList2 = n
	}
	idx := 0
	for n := list1; n != nil; idx++ {
		if idx == a-1 {
			n, n.Next = n.Next, list2
			continue
		}
		if idx == b {
			endList2.Next, n.Next = n.Next, nil
			break
		}
		n = n.Next
	}
	return list1
}
