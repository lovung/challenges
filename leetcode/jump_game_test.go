package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	assert.True(t, canJump([]int{2, 3, 1, 1, 4}))
	assert.False(t, canJump([]int{3, 2, 1, 0, 4}))
	assert.True(t, canJump([]int{1, 2, 3}))
}

func TestJumpGame2(t *testing.T) {
	assert.Equal(t, 2, jumpGame2([]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}))
	assert.Equal(t, 2, jumpGame2([]int{2, 3, 1, 1, 4}))
	assert.Equal(t, 2, jumpGame2([]int{3, 2, 1, 1, 4}))
	assert.Equal(t, 2, jumpGame2([]int{1, 2, 3}))
	assert.Equal(t, 1, jumpGame2([]int{2, 1}))
	assert.Equal(t, 1, jumpGame2([]int{5, 4, 3, 2, 1, 0}))
	assert.Equal(t, 1, jumpGame2([]int{2, 0}))
	assert.Equal(t, 1, jumpGame2([]int{2, 3, 1}))
	assert.Equal(t, 3, jumpGame2([]int{1, 1, 1, 1}))
}
