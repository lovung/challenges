package topinterview150

import (
	"container/heap"

	"github.com/lovung/ds/heaps"
)

func topKFrequent(nums []int, k int) []int {
	counterMap := make(map[int]int)
	for i := range nums {
		counterMap[nums[i]]++
	}
	h := heaps.MaxHeapWithValue[int]{}
	for k, v := range counterMap {
		item := heaps.HeapItem[int]{
			Ref:   v,
			Value: k,
		}
		heap.Push(&h, &item)
	}
	ret := make([]int, 0, k)
	for range k {
		item := heap.Pop(&h).(*heaps.HeapItem[int])
		ret = append(ret, item.Value.(int))
	}
	return ret
}
