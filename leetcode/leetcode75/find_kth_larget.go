package leetcode75

import (
	"container/heap"

	"github.com/lovung/ds/heaps"
)

// https://leetcode.com/problems/kth-largest-element-in-an-array
func findKthLargest(nums []int, k int) int {
	h := heaps.MaxHeap[int]{}
	heap.Init(&h)
	for i := range nums {
		heap.Push(&h, nums[i])
	}
	var ret int
	for i := 0; i < k; i++ {
		ret = heap.Pop(&h).(int)
	}
	return ret
}
