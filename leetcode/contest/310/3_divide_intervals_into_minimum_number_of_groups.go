package contest

import (
	"container/heap"
	"sort"

	"github.com/lovung/ds/heaps"
)

// Link: https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/
func minGroups(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	waitRightsHeap := heaps.MinHeap[int]{}
	heap.Init(&waitRightsHeap)
	for i := range intervals {
		if len(waitRightsHeap) == 0 {
			heap.Push(&waitRightsHeap, intervals[i][1])
		} else {
			waitRight := heap.Pop(&waitRightsHeap).(int)
			if waitRight >= intervals[i][0] {
				heap.Push(&waitRightsHeap, waitRight)
			}
			heap.Push(&waitRightsHeap, intervals[i][1])
		}
	}
	return waitRightsHeap.Len()
}
