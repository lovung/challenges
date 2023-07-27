package grind75

import "github.com/lovung/ds/lists"

func hasCycle(head *lists.ListNode[int]) bool {
	runner := head
	mark := make(map[*lists.ListNode[int]]bool)
	for runner != nil {
		if mark[runner] {
			return true
		}
		mark[runner] = true
		runner = runner.Next
	}
	return false
}
