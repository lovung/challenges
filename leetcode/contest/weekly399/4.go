package weekly399

// https://leetcode.com/problems/maximum-sum-of-subsequence-with-non-adjacent-elements/
func maximumSumSubsequence1(nums []int, queries [][]int) int {
	const mod = 1e9 + 7
	res := 0
	// Declare dp array
	n := len(nums)
	dp := make([][2]int, n)
	for _, q := range queries {
		nums[q[0]] = q[1]
		add := findMaxSum(nums, dp, q[0])
		res += add
		res %= mod
	}
	return res
}

func findMaxSum(arr []int, dp [][2]int, recal int) int {
	n := len(arr)
	if n == 1 {
		return max(0, arr[0])
	}

	// Initialize the values in dp array
	dp[0][0] = 0
	dp[0][1] = arr[0]

	// Loop to find the maximum possible sum
	for i := max(1, recal); i < n; i++ {
		dp[i][1] = dp[i-1][0] + arr[i]
		dp[i][0] = max(dp[i-1][1], dp[i-1][0])
	}

	// Return the maximum sum
	return max(dp[n-1][0], dp[n-1][1])
}
