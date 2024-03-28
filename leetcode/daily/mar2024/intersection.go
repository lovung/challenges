package mar2024

// https://leetcode.com/problems/intersection-of-two-arrays
func intersection(nums1 []int, nums2 []int) []int {
	m1 := make([]bool, 1001)
	m2 := make([]bool, 1001)
	for _, n1 := range nums1 {
		m1[n1] = true
	}
	ret := make([]int, 0, 1000)
	for _, n2 := range nums2 {
		if m1[n2] && !m2[n2] {
			ret = append(ret, n2)
			m2[n2] = true
		}
	}
	return ret
}
