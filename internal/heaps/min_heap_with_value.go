package heaps

import "golang.org/x/exp/constraints"

type HeapItem[T constraints.Ordered] struct {
	Ref   T
	Value any
}

// An MinHeapWithValue is a min-heap of ints.
type MinHeapWithValue[T constraints.Ordered] []HeapItem[T]

func (h MinHeapWithValue[T]) Len() int           { return len(h) }
func (h MinHeapWithValue[T]) Less(i, j int) bool { return h[i].Ref < h[j].Ref }
func (h MinHeapWithValue[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push to heap
func (h *MinHeapWithValue[T]) Push(x any) {
	*h = append(*h, x.(HeapItem[T]))
}

// Pop from heap
func (h *MinHeapWithValue[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
