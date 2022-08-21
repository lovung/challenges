package heaps

import "golang.org/x/exp/constraints"

// An UniqMinHeap is a min-heap of ints.
type UniqMinHeap[T constraints.Ordered] struct {
	data []T
	mark map[T]struct{}
}

func NewUniqMinHeap[T constraints.Ordered]() UniqMinHeap[T] {
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
