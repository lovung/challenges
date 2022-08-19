package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMedianSortedArrays(t *testing.T) {
	t.Run("#1", func(t *testing.T) {
		nums1 := []int{1, 2}
		nums2 := []int{3, 4}
		assert.Equal(t, 2.5, findMedianSortedArrays(nums1, nums2))
	})
	t.Run("#2", func(t *testing.T) {
		nums1 := []int{1, 2, 5}
		nums2 := []int{3, 4}
		assert.Equal(t, 3.0, findMedianSortedArrays(nums1, nums2))
	})
}
