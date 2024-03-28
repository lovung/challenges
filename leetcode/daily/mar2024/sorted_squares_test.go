package mar2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortedSquares(t *testing.T) {
	assert.Equal(t, []int{0, 1, 9, 16, 100}, sortedSquares([]int{-4, -1, 0, 3, 10}))
	assert.Equal(t, []int{4, 9, 9, 49, 121}, sortedSquares([]int{-7, -3, 2, 3, 11}))
}
