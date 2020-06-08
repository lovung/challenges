package junechallenge

func change(amount int, coins []int) int {
	var (
		n  = len(coins)
		dp = make([]int, amount+1)
	)
	dp[0] = 1
	for j := 1; j <= n; j++ {
		for i := coins[j-1]; i <= amount; i++ {
			dp[i] += dp[i-coins[j-1]]
		}
	}
	return dp[amount]
}
