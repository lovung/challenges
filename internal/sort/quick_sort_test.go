package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	nums := []int{0, -1, 2, -3, 4, 5}
	QuickSort(nums)
	assert.Equal(t, []int{-3, -1, 0, 2, 4, 5}, nums)
}
