package mar2024

// https://leetcode.com/problems/minimum-common-value
func getCommon(nums1 []int, nums2 []int) int {
	p, q := 0, 0
	for p < len(nums1) && q < len(nums2) {
		if nums1[p] == nums2[q] {
			return nums1[p]
		}
		if nums1[p] < nums2[q] {
			p++
		} else {
			q++
		}
	}
	return -1
}
