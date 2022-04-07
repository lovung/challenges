package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findRightInternal(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		intervals := [][]int{
			{1, 2},
		}
		assert.Equal(t, []int{-1}, findRightInterval(intervals))
	})
	t.Run("example 2", func(t *testing.T) {
		intervals := [][]int{
			{3, 4},
			{2, 3},
			{1, 2},
		}
		assert.Equal(t, []int{-1, 0, 1}, findRightInterval(intervals))
	})
	t.Run("example 3", func(t *testing.T) {
		intervals := [][]int{
			{1, 4},
			{2, 3},
			{3, 4},
		}
		assert.Equal(t, []int{-1, 2, -1}, findRightInterval(intervals))
	})
}
