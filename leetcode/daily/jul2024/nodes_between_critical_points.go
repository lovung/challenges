package jul2024

import "github.com/lovung/ds/lists"

// https://leetcode.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/
func nodesBetweenCriticalPoints(head *lists.ListNode[int]) []int {
	res := []int{-1, -1}
	prev, node := head, head.Next
	firstIdx, lastIdx, i := -1, -1, 1
	for node != nil {
		if node.Next == nil {
			break
		}
		if (prev.Val > node.Val && node.Val < node.Next.Val) ||
			(prev.Val < node.Val && node.Val > node.Next.Val) {
			if firstIdx == -1 {
				firstIdx, lastIdx = i, i
			} else {
				res[1] = i - firstIdx
				if res[0] == -1 {
					res[0] = i - lastIdx
				} else {
					res[0] = min(res[0], i-lastIdx)
				}
				lastIdx = i
			}
		}
		prev, node = node, node.Next
		i++
	}
	return res
}
