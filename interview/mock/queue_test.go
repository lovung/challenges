package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_circleQueue(t *testing.T) {
	q := New[int](3)
	assert.Equal(t, 0, q.Len())
	v, ok := q.Dequeue()
	assert.False(t, ok)
	assert.True(t, q.Enqueue(1))
	assert.True(t, q.Enqueue(2))
	assert.True(t, q.Enqueue(3))
	assert.False(t, q.Enqueue(4))
	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 2, v)
	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 3, v)
	v, ok = q.Dequeue()
	assert.False(t, ok)

	assert.True(t, q.Enqueue(4))
	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 4, v)
}
