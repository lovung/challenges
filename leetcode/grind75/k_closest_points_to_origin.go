package grind75

import (
	"container/heap"

	"github.com/lovung/ds/heaps"
)

// Link: https://leetcode.com/problems/k-closest-points-to-origin/
func kClosest(points [][]int, k int) [][]int {
	minHeap := heaps.MinHeapWithValue[int]{}
	heap.Init(&minHeap)
	for _, p := range points {
		heap.Push(&minHeap, &heaps.HeapItem[int]{
			Ref:   p[0]*p[0] + p[1]*p[1],
			Value: []int{p[0], p[1]},
		})
	}
	res := make([][]int, 0)
	for i := 0; i < k; i++ {
		nearest := heap.Pop(&minHeap).(*heaps.HeapItem[int])
		res = append(res, nearest.Value.([]int))
	}
	return res
}
