package contest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSum(t *testing.T) {
	assert.Equal(t, 88, maxSum([]int{51, 71, 17, 24, 42}))
	assert.Equal(t, 143, maxSum([]int{51, 71, 17, 24, 42, 72}))
	assert.Equal(t, -1, maxSum([]int{1, 2, 3, 4}))
}
