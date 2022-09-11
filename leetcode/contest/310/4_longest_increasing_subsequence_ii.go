package contest

import "github.com/lovung/ds/maths"

// Link: https://leetcode.com/problems/longest-increasing-subsequence-ii/
func lengthOfLIS(nums []int, k int) int {
	lis := make([]int, 1e5+1)
	for i := range nums {
		if nums[i] >= k {
			lis[nums[i]] = maths.Max(lis[nums[i]-k:nums[i]]...) + 1
		} else {
			lis[nums[i]] = maths.Max(lis[:nums[i]]...) + 1
		}
	}
	return maths.Max(lis...)
}
