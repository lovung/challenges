package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	q.Push(1)
	assert.Equal(t, 1, q.Len())
	assert.Equal(t, 1, q.Pop())
	assert.Equal(t, 0, q.Len())
}
