package medium

import (
	"container/heap"

	"github.com/lovung/ds/heaps"
)

// Link: https://leetcode.com/problems/smallest-number-in-infinite-set/

// SmallestInfiniteSet object will hold the set
type SmallestInfiniteSet struct {
	minHeap *heaps.UniqMinHeap[int]
}

func Constructor() SmallestInfiniteSet {
	h := heaps.NewUniqMinHeap[int]()
	heap.Init(&h)
	for i := 1; i <= 1000; i++ {
		heap.Push(&h, i)
	}
	return SmallestInfiniteSet{
		minHeap: &h,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	return heap.Pop(this.minHeap).(int)
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	heap.Push(this.minHeap, num)
}
