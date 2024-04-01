package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMedianSortedArrays(t *testing.T) {
	t.Run("#1", func(t *testing.T) {
		nums1 := []int{1, 2}
		nums2 := []int{3, 4}
		assert.EqualValues(t, 2.5, findMedianSortedArrays(nums1, nums2))
	})
	t.Run("#2", func(t *testing.T) {
		nums1 := []int{1, 2, 5}
		nums2 := []int{3, 4}
		assert.EqualValues(t, 3.0, findMedianSortedArrays(nums1, nums2))
	})
	t.Run("#3", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 3, 4}
		nums2 := []int{2, 5, 5, 6}
		assert.EqualValues(t, 3, findMedianSortedArrays(nums1, nums2))
	})
	t.Run("#4", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 4, 5}
		nums2 := []int{6}
		assert.EqualValues(t, 3.5, findMedianSortedArrays(nums1, nums2))
	})
	t.Run("#5", func(t *testing.T) {
		nums1 := []int{1e5}
		nums2 := []int{1e5 + 1}
		assert.EqualValues(t, 1e5+0.5, findMedianSortedArrays(nums1, nums2))
	})
}

func Test_findMedianSortedArrays2(t *testing.T) {
	t.Run("#1", func(t *testing.T) {
		nums1 := []int{1, 2}
		nums2 := []int{3, 4}
		assert.EqualValues(t, 2.5, findMedianSortedArrays2(nums1, nums2))
	})
	t.Run("#2", func(t *testing.T) {
		nums1 := []int{1, 2, 5}
		nums2 := []int{3, 4}
		assert.EqualValues(t, 3.0, findMedianSortedArrays(nums1, nums2))
	})
	t.Run("#3", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 3, 4}
		nums2 := []int{2, 5, 5, 6}
		assert.EqualValues(t, 3, findMedianSortedArrays2(nums1, nums2))
	})
	t.Run("#4", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 4, 5}
		nums2 := []int{6}
		assert.EqualValues(t, 3.5, findMedianSortedArrays2(nums1, nums2))
	})
	t.Run("#5", func(t *testing.T) {
		nums1 := []int{1e5}
		nums2 := []int{1e5 + 1}
		assert.EqualValues(t, 1e5+0.5, findMedianSortedArrays2(nums1, nums2))
	})
}
