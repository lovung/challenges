package problems

import "math"

// Link: https://leetcode.com/problems/minimum-average-difference
func minimumAverageDifference(nums []int) int {
	n := len(nums)
	sums := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		sums[i] = sum
	}
	smallestIdx := -1
	smallest := math.MaxInt
	for i := 0; i < n-1; i++ {
		x := absDiff(sums[i]/(i+1), (sum-sums[i])/(n-i-1))
		if smallest > x {
			smallest = x
			smallestIdx = i
		}
	}
	x := absDiff(sums[n-1]/n, 0)
	if smallest > x {
		smallest = x
		smallestIdx = n - 1
	}
	return smallestIdx
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
