package hard

import (
	"container/heap"
	"sort"

	"github.com/lovung/ds/heaps"
)

// Link: https://leetcode.com/problems/maximum-performance-of-a-team/
type engineer struct {
	speed      int
	efficiency int
}

const mod = 1e9 + 7

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	engineers := make([]*engineer, 0, n)
	for i := range speed {
		engineers = append(engineers, &engineer{
			speed:      speed[i],
			efficiency: efficiency[i],
		})
	}
	sort.Slice(engineers, func(i, j int) bool {
		return engineers[i].efficiency > engineers[j].efficiency
	})
	totalSpeed := 0
	maxPerf := 0
	minHeap := heaps.MinHeap[int]{}
	heap.Init(&minHeap)
	for i := range engineers {
		totalSpeed += engineers[i].speed
		heap.Push(&minHeap, engineers[i].speed)
		if i >= k && minHeap.Len() > 0 {
			minSpeedInTeam := heap.Pop(&minHeap).(int)
			totalSpeed -= minSpeedInTeam
		}
		maxPerf = max(maxPerf, totalSpeed*engineers[i].efficiency)
	}
	return maxPerf % mod
}
