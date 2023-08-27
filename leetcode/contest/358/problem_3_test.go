package contest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minAbsoluteDifference(t *testing.T) {
	assert.Equal(t, 0, minAbsoluteDifference([]int{4, 3, 2, 4}, 2))
	assert.Equal(t, 1, minAbsoluteDifference([]int{5, 3, 2, 10, 15}, 1))
	assert.Equal(t, 3, minAbsoluteDifference([]int{1, 2, 3, 4}, 3))
}
