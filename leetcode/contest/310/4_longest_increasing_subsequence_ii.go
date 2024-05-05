package contest

import (
	"github.com/lovung/ds/maths"
	"github.com/lovung/ds/trees"
)

// Link: https://leetcode.com/problems/longest-increasing-subsequence-ii/
func lengthOfLIS(nums []int, k int) int {
	lis := make([]int, 1e5+1)
	ans := 0
	for i := range nums {
		l := max(0, nums[i]-k)
		lis[nums[i]] = maths.Max(lis[l:nums[i]]...) + 1
		ans = max(ans, lis[nums[i]])
	}
	return ans
}

func lengthOfLIS2(nums []int, k int) int {
	ans := 0
	seg := trees.NewSegmentTreeWithArray(1e5+1, maxFn)
	for i := range nums {
		l := max(0, nums[i]-k)
		r := nums[i]
		res := seg.Query(l, r) + 1
		ans = max(res, ans)
		seg.Update(nums[i], res)
	}
	return ans
}

func maxFn(a, b int) int {
	return max(a, b)
}
