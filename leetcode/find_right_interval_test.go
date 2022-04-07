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
	t.Run("example 4", func(t *testing.T) {
		// [[4,5],[2,3],[1,2]]
		intervals := [][]int{
			{4, 5},
			{2, 3},
			{1, 2},
		}
		assert.Equal(t, []int{-1, 0, 1}, findRightInterval(intervals))
	})
	t.Run("example 5", func(t *testing.T) {
		// input: [[1,12],[2,9],[3,10],[13,14],[15,16],[16,17]]
		// expect: [3,3,3,4,5,-1]
		intervals := [][]int{
			{1, 12},
			{2, 9},
			{3, 10},
			{13, 14},
			{15, 16},
			{16, 17},
		}
		assert.Equal(t, []int{3, 3, 3, 4, 5, -1}, findRightInterval(intervals))
	})
}

func Test_binarySearchForIndexedInterval(t *testing.T) {
	cmpFunc := func(val int, slice []*indexedInterval) func(i int) int {
		return func(i int) int {
			if val < slice[i].interval[0] {
				return -1
			} else if val > slice[i].interval[0] {
				return 1
			}
			return 0
		}
	}
	t.Run("1", func(t *testing.T) {
		slice := []*indexedInterval{
			{index: 0, interval: []int{1, 2}},
			{index: 1, interval: []int{2, 3}},
			{index: 2, interval: []int{4, 6}},
			{index: 3, interval: []int{7, 10}},
			{index: 4, interval: []int{8, 9}},
			{index: 5, interval: []int{12, 15}},
			{index: 6, interval: []int{15, 15}},
		}

		assert.Equal(t, 0, binarySearchForIndexedInterval(slice, cmpFunc(1, slice)))
		assert.Equal(t, 1, binarySearchForIndexedInterval(slice, cmpFunc(2, slice)))
		assert.Equal(t, 2, binarySearchForIndexedInterval(slice, cmpFunc(3, slice)))
		assert.Equal(t, 2, binarySearchForIndexedInterval(slice, cmpFunc(4, slice)))
		assert.Equal(t, 6, binarySearchForIndexedInterval(slice, cmpFunc(15, slice)))
		assert.Equal(t, -1, binarySearchForIndexedInterval(slice, cmpFunc(16, slice)))
	})
	t.Run("2", func(t *testing.T) {
		slice := []*indexedInterval{
			{index: 2, interval: []int{1, 2}},
			{index: 1, interval: []int{2, 3}},
			{index: 0, interval: []int{3, 4}},
		}
		assert.Equal(t, 2, binarySearchForIndexedInterval(slice, cmpFunc(3, slice)))
		assert.Equal(t, -1, binarySearchForIndexedInterval(slice, cmpFunc(4, slice)))
	})
	t.Run("3", func(t *testing.T) {
		slice := []*indexedInterval{
			{index: 2, interval: []int{1, 2}},
			{index: 1, interval: []int{2, 3}},
			{index: 0, interval: []int{4, 5}},
		}
		assert.Equal(t, 2, binarySearchForIndexedInterval(slice, cmpFunc(3, slice)))
		assert.Equal(t, -1, binarySearchForIndexedInterval(slice, cmpFunc(5, slice)))
	})
}
