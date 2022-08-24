package heaps

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeapWithValue(t *testing.T) {
	var h MinHeapWithValue[int]
	heap.Init(&h)
	heap.Push(&h, &HeapItem[int]{Ref: 3, Value: 3})
	heap.Push(&h, &HeapItem[int]{Ref: 1, Value: 1})
	heap.Push(&h, &HeapItem[int]{Ref: 2, Value: 2})
	assert.Equal(t, 3, h.Len())
	assert.Equal(t, 1, heap.Pop(&h).(*HeapItem[int]).Value.(int))
	assert.Equal(t, 2, heap.Pop(&h).(*HeapItem[int]).Ref)
	assert.Equal(t, 1, h.Len())
}
