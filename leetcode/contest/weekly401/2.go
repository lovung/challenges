package weekly401

func valueAfterKSeconds(n int, k int) int {
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for range k {
		for i := 1; i < n; i++ {
			dp[i] += dp[i-1]
			dp[i] %= 1e9 + 7
		}
	}
	return dp[n-1]
}
