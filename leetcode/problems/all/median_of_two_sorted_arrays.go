package problems

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	N := m + n
	mergedNums := append(nums1, nums2...)
	sort.SliceStable(mergedNums, func(i, j int) bool {
		return mergedNums[i] < mergedNums[j]
	})
	if N&0x01 == 0 {
		return float64((mergedNums[N/2-1] + mergedNums[N/2])) / 2
	}
	return float64(mergedNums[N/2])
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	nums1 = append(nums1, 1e6+1)
	nums2 = append(nums2, 1e6+1)
	p, q := -1, -1
	val := 0
	for range (n + 1) / 2 {
		if nums1[p+1] < nums2[q+1] {
			p++
			val = nums1[p]
		} else {
			q++
			val = nums2[q]
		}
	}
	if n%2 != 0 {
		return float64(val)
	}
	return float64(val+min(nums1[p+1], nums2[q+1])) / 2
}
