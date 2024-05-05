package maxprofit

// Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
func maxProfitIII(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	k := 2
	dp := make([][]int, k+1)
	dp[0] = make([]int, len(prices))
	for i := 1; i <= k; i++ {
		dp[i] = make([]int, len(prices))
		localMax := dp[i-1][0] - prices[0]
		for j := 1; j < len(prices); j++ {
			dp[i][j] = max(dp[i][j-1], prices[j]+localMax)
			localMax = max(localMax, dp[i-1][j]-prices[j])
		}
	}
	return dp[k][len(prices)-1]
}
