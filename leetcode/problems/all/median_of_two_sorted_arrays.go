package problems

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	N := m + n
	mergedNums := make([]int, N)
	mergedNums = append(nums1, nums2...)
	sort.SliceStable(mergedNums, func(i, j int) bool {
		return mergedNums[i] < mergedNums[j]
	})
	if N&0x01 == 0 {
		return float64((mergedNums[N/2-1] + mergedNums[N/2])) / 2
	}
	return float64(mergedNums[N/2])
}
