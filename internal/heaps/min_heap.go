package heaps

import "golang.org/x/exp/constraints"

// MinHeap is a min-heap of ints.
type MinHeap[T constraints.Ordered] []T

func (h MinHeap[T]) Len() int           { return len(h) }
func (h MinHeap[T]) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push to heap
func (h *MinHeap[T]) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(T))
}

// Pop from heap
func (h *MinHeap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
