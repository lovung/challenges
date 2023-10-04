package topinterview150

// Link: https://leetcode.com/problems/merge-sorted-array
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j := 0, 0; j < n; i++ {
		if i >= m {
			nums1[i] = nums2[j]
			j++
			continue
		}
		if nums1[i] > nums2[j] {
			copy(nums1[i+1:m+1], nums1[i:m])
			nums1[i] = nums2[j]
			m++
			j++
		}
	}
}
