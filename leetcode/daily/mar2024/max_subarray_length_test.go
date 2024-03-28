package mar2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubarrayLength(t *testing.T) {
	assert.Equal(t, 6, maxSubarrayLength([]int{1, 2, 3, 1, 2, 3, 1, 2}, 2))
	assert.Equal(t, 2, maxSubarrayLength([]int{1, 2, 1, 2, 1, 2, 1, 2}, 1))
	assert.Equal(t, 4, maxSubarrayLength([]int{5, 5, 5, 5, 5, 5, 5, 5}, 4))
	assert.Equal(t, 2, maxSubarrayLength([]int{1, 4, 4, 3}, 1))
}
