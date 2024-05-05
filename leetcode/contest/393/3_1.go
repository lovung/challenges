package contest393

import (
	"container/heap"

	"github.com/lovung/ds/heaps"
)

type combine struct {
	val  int
	deno int
}

func findKthSmallest_1(coins []int, k int) int64 {
	h := heaps.MinHeapWithValue[int]{}
	heap.Init(&h)
	for i := range coins {
		item := heaps.HeapItem[int]{
			Ref:   coins[i],
			Value: combine{coins[i], coins[i]},
		}
		heap.Push(&h, &item)
	}
	lcm := int64(LCM(coins...))
	res := int64(0)
	arr := make([]int64, 0, 100)
	for _k := k; _k > 0; {
		itemPop := heap.Pop(&h).(*heaps.HeapItem[int])
		itemPopCom := itemPop.Value.(combine)
		if int64(itemPopCom.val) > res {
			res = int64(itemPopCom.val)
			arr = append(arr, res)
			_k--
			if _k == 0 {
				return res
			}
			if res == lcm {
				break
			}
		}
		newVal := itemPopCom.val + itemPopCom.deno
		item := heaps.HeapItem[int]{
			Ref:   newVal,
			Value: combine{newVal, itemPopCom.deno},
		}
		heap.Push(&h, &item)
	}
	milestone := int64((k-1)/len(arr)) * res
	index := (k - 1) % len(arr)
	return milestone + arr[index]
}

func isLCM(n int64, coins []int) bool {
	for i := range coins {
		if n%int64(coins[i]) != 0 {
			return false
		}
	}
	return true
}
