package leetcode75

import "github.com/lovung/ds/maths"

// https://leetcode.com/problems/find-pivot-index
func pivotIndex(nums []int) int {
	sum := maths.Sum(nums...)
	curSum := 0
	for i := range nums {
		if sum-nums[i]-curSum == curSum {
			return i
		}
		curSum += nums[i]
	}
	return -1
}
