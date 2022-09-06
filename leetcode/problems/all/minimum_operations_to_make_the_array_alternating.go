package problems

import (
	"container/heap"

	"github.com/lovung/ds/heaps"
	"github.com/lovung/ds/maths"
)

// Link: https://leetcode.com/problems/minimum-operations-to-make-the-array-alternating/
func minimumOperations(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	oddCntMap := make(map[int]int)
	evenCntMap := make(map[int]int)
	for i := range nums {
		if i%2 == 0 {
			evenCntMap[nums[i]]++
		} else {
			oddCntMap[nums[i]]++
		}
	}
	oddHeap, evenHeap := heaps.MaxHeapWithValue[int]{}, heaps.MaxHeapWithValue[int]{}
	heap.Init(&oddHeap)
	heap.Init(&evenHeap)
	for k, v := range oddCntMap {
		item := heaps.HeapItem[int]{
			Ref:   v,
			Value: k,
		}
		heap.Push(&oddHeap, &item)
	}
	for k, v := range evenCntMap {
		item := heaps.HeapItem[int]{
			Ref:   v,
			Value: k,
		}
		heap.Push(&evenHeap, &item)
	}
	firstOdd := heap.Pop(&oddHeap).(*heaps.HeapItem[int])
	firstEven := heap.Pop(&evenHeap).(*heaps.HeapItem[int])
	if firstOdd.Value.(int) != firstEven.Value.(int) { // nums[i - 1] != nums[i]
		return len(nums) - firstOdd.Ref - firstEven.Ref
	}
	var secondOdd, secondEven *heaps.HeapItem[int]
	if oddHeap.Len() > 0 {
		secondOdd = heap.Pop(&oddHeap).(*heaps.HeapItem[int])
	}
	if evenHeap.Len() > 0 {
		secondEven = heap.Pop(&evenHeap).(*heaps.HeapItem[int])
	}
	switch {
	case secondOdd == nil && secondEven == nil:
		return maths.Min(firstOdd.Ref, firstEven.Ref)
	case secondOdd == nil && secondEven != nil:
		return len(nums) - firstOdd.Ref - secondEven.Ref
	case secondOdd != nil && secondEven == nil:
		return len(nums) - firstEven.Ref - secondOdd.Ref
	default:
		if secondOdd.Ref+firstEven.Ref > secondEven.Ref+firstOdd.Ref {
			return len(nums) - secondOdd.Ref - firstEven.Ref
		}
		return len(nums) - secondEven.Ref - firstOdd.Ref
	}
}
