package aug2023

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchMatrix(t *testing.T) {
	assert.True(t, searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	assert.False(t, searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
	assert.True(t, searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 10))
	assert.True(t, searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 23))
	assert.True(t, searchMatrix([][]int{{1}}, 1))
	assert.True(t, searchMatrix([][]int{{1}, {3}}, 3))
	assert.False(t, searchMatrix([][]int{{1}, {3}}, 2))
	assert.False(t, searchMatrix([][]int{{2}, {3}}, 1))
}
