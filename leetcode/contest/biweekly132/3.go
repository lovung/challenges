package biweekly132

import (
	"container/heap"
)

type node struct{ len, rk int }

// https://leetcode.com/problems/find-the-maximum-length-of-a-good-subsequence-i/description/
// Wrong anwser
func maximumLength_wrong(nums []int, k int) int {
	res := 0

	maxHeap := MaxHeapWithValue[int]{}
	heap.Init(&maxHeap)
	dp := make(map[int][]*node)
	dp[nums[0]] = append(dp[nums[0]], &node{1, 0})
	heap.Push(&maxHeap, &HeapItem[int]{
		Ref:   1,
		Value: nums[0],
	})
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if len(dp[num]) == 0 {
			dp[num] = append(dp[num], &node{1, 0})
		} else {
			for i := range dp[num] {
				dp[num][i].len++
			}
		}
		maxValue := heap.Pop(&maxHeap).(*HeapItem[int])
		if maxValue.Value == num && len(maxHeap) > 0 {
			secondValue := heap.Pop(&maxHeap).(*HeapItem[int])
			if secondValue.Value == num && len(maxHeap) > 0 {
				thirdValue := heap.Pop(&maxHeap).(*HeapItem[int])
				secondValue = thirdValue
			}
			maxValue = secondValue
		}
		if dp[num][0].len < maxValue.Ref+1 {
			if n := getMaxValueLessThanK(dp[maxValue.Value], k); n != nil {
				if len(dp[num]) == 1 {
					dp[num] = append(dp[num], &node{n.len + 1, n.rk + 1})
				} else {
					dp[num][1] = &node{n.len + 1, n.rk + 1}
				}
				res = max(res, dp[num][1].len)
			}
		}
		heap.Push(&maxHeap, maxValue)
		for _, n := range dp[num] {
			heap.Push(&maxHeap, &HeapItem[int]{
				Ref:   n.len,
				Value: num,
			})
		}
		res = max(res, dp[num][0].len)
	}
	return res
}

func getMaxValueLessThanK(nodes []*node, k int) *node {
	var res *node
	for _, n := range nodes {
		if n.rk < k && (res == nil || n.len > res.len) {
			res = n
		}
	}
	return res
}

type HeapItem[T ~int] struct {
	Ref   T
	Value T

	index int
}

// An MaxHeapWithValue is a max-heap of ints.
type MaxHeapWithValue[T ~int] []*HeapItem[T]

func (h MaxHeapWithValue[T]) Len() int           { return len(h) }
func (h MaxHeapWithValue[T]) Less(i, j int) bool { return h[i].Ref > h[j].Ref }
func (h MaxHeapWithValue[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

// Push to heap
func (h *MaxHeapWithValue[T]) Push(x any) {
	n := len(*h)
	item := x.(*HeapItem[T])
	item.index = n
	*h = append(*h, item)
}

// Pop from heap
func (h *MaxHeapWithValue[T]) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}
