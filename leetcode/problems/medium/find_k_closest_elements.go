package medium

import (
	"container/heap"
	"sort"

	"github.com/lovung/ds/heaps"
)

// Link: https://leetcode.com/problems/find-k-closest-elements/
func findClosestElements(arr []int, k int, x int) []int {
	l, r := binarySearch(arr, x)
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		// item := heap.Pop(&minHeap).(*heaps.HeapItem[int])
		if l < 0 {
			res = append(res, arr[r])
			r++
			continue
		}
		if r >= len(arr) {
			res = append(res, arr[l])
			l--
			continue
		}
		if abs(arr[l]-x) <= abs(arr[r]-x) {
			res = append(res, arr[l])
			l--
		} else {
			res = append(res, arr[r])
			r++
		}
	}
	sort.Ints(res)
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 1 , 3 , 5, 7
// 0 -> -1, 0
// 1 -> 0, 1
// 2 -> 0, 1
// 3 -> 1, 2
// 4 -> 1, 2
// 5 -> 2, 3
// 6 -> 2, 3
// 7 -> 3, 4
// 8 -> 3, 4
func binarySearch(arr []int, x int) (int, int) {
	l, r := 0, len(arr)-1
	if x < arr[l] {
		return -1, 0
	}
	if x > arr[r] {
		return len(arr) - 1, len(arr)
	}
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == x {
			return mid, mid + 1
		}
		if arr[mid] < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r, l
}

func findClosestElements2(arr []int, k int, x int) []int {
	minHeap := heaps.MinHeapWithValue[int]{}
	heap.Init(&minHeap)
	for i := range arr {
		heap.Push(&minHeap, &heaps.HeapItem[int]{
			Ref:   calculateOrderValue(arr[i] - x),
			Value: arr[i],
		})
	}
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		item := heap.Pop(&minHeap).(*heaps.HeapItem[int])
		res = append(res, item.Value.(int))
	}
	sort.Ints(res)
	return res
}

func calculateOrderValue(x int) int {
	if x < 0 {
		return -2 * x
	}
	return 2*x + 1
}
