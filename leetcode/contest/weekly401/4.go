package weekly401

import (
	"container/heap"
	"slices"
	"sort"
)

// https://leetcode.com/problems/maximum-total-reward-using-operations-ii/description/
func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	n := len(rewardValues)
	dp := [50000]int{}
	for i, v := range rewardValues {
		if i == 0 || rewardValues[i-1] != v {
			lim := min(v, rewardValues[n-1]-v)
			for x := 0; x < lim; x++ {
				dp[v+dp[x]] = v + dp[x]
			}
		}
	}
	return rewardValues[n-1] + slices.Max(dp[0:rewardValues[n-1]])
}

// TLE
func maxTotalReward2(rewardValues []int) int {
	sort.Ints(rewardValues)
	h := NewUniqMinHeap[int]()
	heap.Init(&h)
	heap.Push(&h, 0)
	res := 0
	for _, v := range rewardValues {
		temp := make([]int, 0, v)
		top := heap.Pop(&h).(int)
		for top < v {
			temp = append(temp, top)
			temp = append(temp, top+v)
			res = max(res, top+v)
			if h.Len() == 0 {
				break
			}
			top = heap.Pop(&h).(int)
		}
		temp = append(temp, top)
		for i := range temp {
			heap.Push(&h, temp[i])
		}
	}
	return res
}

// An UniqMinHeap is a min-heap of ints.
type UniqMinHeap[T ~int] struct {
	data []T
	mark map[T]struct{}
}

func NewUniqMinHeap[T ~int]() UniqMinHeap[T] {
	return UniqMinHeap[T]{
		data: make([]T, 0),
		mark: make(map[T]struct{}),
	}
}

func (h UniqMinHeap[T]) Len() int           { return len(h.data) }
func (h UniqMinHeap[T]) Less(i, j int) bool { return h.data[i] < h.data[j] }
func (h UniqMinHeap[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

// Push to heap
func (h *UniqMinHeap[T]) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	if _, ok := h.mark[x.(T)]; ok {
		return
	}
	h.data = append(h.data, x.(T))
	h.mark[x.(T)] = struct{}{}
}

// Pop from heap
func (h *UniqMinHeap[T]) Pop() any {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	delete(h.mark, x)
	return x
}
