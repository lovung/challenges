package contest

import "github.com/lovung/ds/maths"

// Link: https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/
func longestSubarray(nums []int) int {
	k := maths.Max(nums...)
	cnt := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			j := i
			for j < len(nums) && nums[j] == k {
				j++
			}
			cnt = max(cnt, j-i)
			i = j - 1
		}
	}
	return cnt
}
