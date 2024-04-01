package mar2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchRange(t *testing.T) {
	assert.Equal(t, []int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	assert.Equal(t, []int{1, 2}, searchRange([]int{5, 7, 7, 8, 8, 10}, 7))
	assert.Equal(t, []int{1, 3}, searchRange([]int{5, 7, 7, 7, 8, 8, 10}, 7))
	assert.Equal(t, []int{1, 6}, searchRange([]int{5, 7, 7, 7, 7, 7, 7, 8, 8, 10}, 7))
	assert.Equal(t, []int{3, 4}, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	assert.Equal(t, []int{-1, -1}, searchRange([]int{}, 0))
}
