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

func jumpGame2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	maxReachIdx := nums[0]
	nextMaxIdx := maxReachIdx
	jumpCount := 1
	for i := 1; i <= maxReachIdx && i < n-1; i++ {
		if i <= nextMaxIdx {
			if maxReachIdx < i+nums[i] {
				maxReachIdx = i + nums[i]
			}
		}
		if i == nextMaxIdx {
			// fmt.Println(i, jumpCount, maxReachIdx)
			jumpCount++
			nextMaxIdx = maxReachIdx
		}
	}
	return jumpCount
}
