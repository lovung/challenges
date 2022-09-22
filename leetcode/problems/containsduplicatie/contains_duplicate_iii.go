package containsduplicatie

import (
	"github.com/lovung/ds/trees"
)

// Link: https://leetcode.com/problems/contains-duplicate-iii/
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	const offset = 1e9 + 1
	const idxOffset = 1
	lastIdx := trees.NewSegmentTreeWithMap[int](2*1e9+1, max)
	for i := range nums {
		if idx := lastIdx.Query(
			nums[i]+offset-valueDiff,
			nums[i]+offset+valueDiff+1,
		); idx > 0 && i-(idx-idxOffset) <= indexDiff {
			return true
		}
		lastIdx.Update(nums[i]+offset, i+idxOffset)
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
