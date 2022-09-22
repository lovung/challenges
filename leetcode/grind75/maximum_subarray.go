package grind75

// Link: https://leetcode.com/problems/maximum-subarray
func maxSubArray(nums []int) int {
	sums := make([]int, len(nums)+1)
	sum := 0
	for i := range nums {
		sum += nums[i]
		sums[i+1] = sum
	}
	maxDiff := sums[len(sums)-1]
	maxRight := sums[len(sums)-1]
	for i := len(sums) - 2; i >= 0; i-- {
		diff := maxRight - sums[i]
		if diff > maxDiff {
			maxDiff = diff
		}
		if sums[i] > maxRight {
			maxRight = sums[i]
		}
	}
	return maxDiff
}
