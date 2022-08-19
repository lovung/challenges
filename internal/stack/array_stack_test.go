package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayStack(t *testing.T) {
	s := NewArrayStack[int](3)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, 3, s.Len())
	assert.Equal(t, []int{1, 2, 3}, s.All())
	assert.Equal(t, 3, s.Peek())
	assert.Equal(t, 3, s.Pop())
	assert.Equal(t, 2, s.Len())
	assert.Equal(t, []int{1, 2}, s.All())
	assert.Equal(t, 2, s.Peek())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 1, s.Len())
	assert.Equal(t, []int{1}, s.All())
	assert.Equal(t, 1, s.Peek())
	assert.Equal(t, 1, s.Pop())
	assert.Equal(t, 0, s.Len())
	assert.Equal(t, []int{}, s.All())
	assert.Equal(t, 0, s.Peek())
	assert.Equal(t, 0, s.Pop())
	s.Push(1)
	assert.Equal(t, 1, s.Len())
	s.Clear()
	assert.Equal(t, 0, s.Len())
}
