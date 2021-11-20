package leetcode

func canJump(nums []int) bool {
	n := len(nums)
	maxReachIdx := nums[0]
	for i := 1; i <= maxReachIdx && i < n; i++ {
		if maxReachIdx < i+nums[i] {
			maxReachIdx = i + nums[i]
		}
	}
	return maxReachIdx >= n-1
}
