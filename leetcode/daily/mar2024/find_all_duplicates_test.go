package mar2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findAllDuplicates(t *testing.T) {
	assert.Equal(t, []int{2, 3}, findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	assert.Equal(t, []int{1}, findDuplicates([]int{1, 1, 2}))
	assert.Equal(t, []int{}, findDuplicates([]int{1}))
}
