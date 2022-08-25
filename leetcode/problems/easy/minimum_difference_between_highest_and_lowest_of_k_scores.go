package easy

import (
	"sort"
)

// Link: https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	min := 100001
	for i := 0; i <= len(nums)-k; i++ {
		if min > nums[i+k-1]-nums[i] {
			min = nums[i+k-1] - nums[i]
		}
	}
	return min
}
