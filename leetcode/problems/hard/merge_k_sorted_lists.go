package hard

import (
	"container/heap"

	"github.com/lovung/ds/heaps"
	. "github.com/lovung/ds/lists"
)

// Link: https://leetcode.com/problems/merge-k-sorted-lists/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode[int]) *ListNode[int] {
	// create a heap with ref is value of linked-list node
	// the key will be the pointer of the node
	minHeap := heaps.MinHeapWithValue[int]{}
	heap.Init(&minHeap)
	// first round, check all the first item of all link lists
	// put them inside the heap
	for i := range lists {
		if lists[i] == nil {
			continue
		}
		heap.Push(&minHeap, &heaps.HeapItem[int]{
			Ref:   lists[i].Val,
			Value: lists[i],
		})
		lists[i] = lists[i].Next
	}
	// create new merged sorted linked list
	var newList *ListNode[int]
	var runner *ListNode[int]

	for minHeap.Len() > 0 {
		smallestItem := heap.Pop(&minHeap).(*heaps.HeapItem[int])
		if runner == nil {
			runner = smallestItem.Value.(*ListNode[int])
			newList = runner
		} else {
			runner.Next = smallestItem.Value.(*ListNode[int])
			runner = runner.Next
		}
		if runner.Next != nil {
			heap.Push(&minHeap, &heaps.HeapItem[int]{
				Ref:   runner.Next.Val,
				Value: runner.Next,
			})
		}
	}
	return newList
}
