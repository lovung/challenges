package medium

// Link: https://leetcode.com/problems/maximum-length-of-repeated-subarray/
func findLength(nums1 []int, nums2 []int) int {
	res := 0
	dp := make([][]int, len(nums1)+1)
	dp[len(nums1)] = make([]int, len(nums2)+1)
	for i := len(nums1) - 1; i >= 0; i-- {
		dp[i] = make([]int, len(nums2)+1)
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				if res < dp[i][j] {
					res = dp[i][j]
				}
			}
		}
	}
	return res
}
