package heaps

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaps(t *testing.T) {
	var h MinHeap[int]
	heap.Init(&h)
	heap.Push(&h, 1)
	heap.Push(&h, 2)
	heap.Push(&h, 3)
	heap.Push(&h, 4)
	assert.Equal(t, 4, h.Len())
	assert.Equal(t, 1, heap.Pop(&h))
	assert.Equal(t, 2, heap.Pop(&h))
	assert.Equal(t, 3, heap.Pop(&h))
	assert.Equal(t, 4, heap.Pop(&h))
	assert.Equal(t, 0, h.Len())
}
