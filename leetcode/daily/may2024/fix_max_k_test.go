package may2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMaxK(t *testing.T) {
	assert.Equal(t, 3, findMaxK([]int{-1, 2, -3, 3}))
	assert.Equal(t, 7, findMaxK([]int{-1, 10, 6, 7, -7, 1}))
	assert.Equal(t, -1, findMaxK([]int{-10, 8, 6, 7, -2, -3}))
}
