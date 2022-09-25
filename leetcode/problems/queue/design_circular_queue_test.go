package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MyCircularQueue(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		q := Constructor(3)
		assert.True(t, q.EnQueue(1))
		assert.True(t, q.EnQueue(2))
		assert.True(t, q.EnQueue(3))
		assert.False(t, q.EnQueue(4))
		assert.Equal(t, 3, q.Rear())
		assert.True(t, q.IsFull())
		assert.True(t, q.DeQueue())
		assert.True(t, q.EnQueue(4))
		assert.Equal(t, 4, q.Rear())
	})
	t.Run("2", func(t *testing.T) {
		q := Constructor(6)
		assert.True(t, q.EnQueue(6))
		assert.Equal(t, 6, q.Rear())
		assert.Equal(t, 6, q.Rear())
		assert.True(t, q.DeQueue())
		assert.True(t, q.EnQueue(5))
		assert.Equal(t, 5, q.Rear())
		assert.True(t, q.DeQueue())
		assert.Equal(t, -1, q.Front())
		assert.False(t, q.DeQueue())
		assert.False(t, q.DeQueue())
		assert.False(t, q.DeQueue())
	})
	t.Run("3", func(t *testing.T) {
		q := Constructor(3)
		assert.Equal(t, -1, q.Rear())
		assert.True(t, q.EnQueue(1))
		assert.True(t, q.DeQueue())
		assert.True(t, q.EnQueue(2))
		assert.True(t, q.DeQueue())
		assert.True(t, q.EnQueue(3))
		assert.True(t, q.DeQueue())
		assert.True(t, q.EnQueue(4))
		assert.Equal(t, 4, q.Rear())
		assert.Equal(t, 4, q.Front())
	})
}
