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
