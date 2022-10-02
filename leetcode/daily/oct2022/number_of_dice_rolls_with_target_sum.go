package oct2022

import "github.com/lovung/ds/maths"

// Link: https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/
const mod = 1e9 + 7

func numRollsToTarget(n int, k int, target int) int {
	dp := make([][]int, n)
	// i+1 is the current n
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target)
	}
	// j+1 is the current target
	for j := 0; j < maths.Min(k, target); j++ {
		dp[0][j] = 1
	}
	for i := 1; i < n; i++ {
		for j := i; j < target; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			if j > k {
				dp[i][j] -= dp[i-1][j-k-1]
			}
			dp[i][j] %= mod
			if dp[i][j] < 0 {
				dp[i][j] += mod
			}
		}
	}
	return dp[n-1][target-1]
}
