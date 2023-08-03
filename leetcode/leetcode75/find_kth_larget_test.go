package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findKthLargest(t *testing.T) {
	assert.Equal(t, 5, findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	assert.Equal(t, 4, findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
