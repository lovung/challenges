package medium

import (
	"container/heap"
	"sort"

	"github.com/lovung/ds/heaps"
)

// Link: https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/
// Burstforce solution
func maxEvents(events [][]int) int {
	if len(events) == 0 {
		return 0
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	h := &heaps.MinHeap[int]{}
	heap.Init(h)
	i, n, count := 0, len(events), 0
	for d := 0; d < 100000; d++ {
		// delete all events that are in the past
		for h.Len() > 0 {
			v := heap.Pop(h)
			if v.(int) < d {
				continue
			} else {
				heap.Push(h, v)
				break
			}
		}
		// insert events (their deadlines) that start at day d
		for i < n && events[i][0] == d {
			heap.Push(h, events[i][1])
			i++
		}
		// schedule the event with least deadline on that day
		if h.Len() > 0 {
			heap.Pop(h)
			count++
		}
		// if no events left, return
		if h.Len() == 0 && i == n {
			break
		}
	}
	return count
}
