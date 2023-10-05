package topinterview150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minSubArrayLen(t *testing.T) {
	assert.Equal(t, 2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	assert.Equal(t, 1, minSubArrayLen(4, []int{1, 4, 4}))
	assert.Equal(t, 0, minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}
