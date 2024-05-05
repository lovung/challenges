package jul2023

// https://leetcode.com/problems/strange-printer
func strangePrinter(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for length := 1; length <= n; length++ {
		for l := 0; l <= n-length; l++ {
			r := l + length - 1
			j := -1
			dp[l][r] = n
			for i := l; i < r; i++ {
				if s[i] != s[r] && j == -1 {
					j = i
				}
				if j != -1 {
					dp[l][r] = min(dp[l][r], 1+dp[j][i]+dp[i+1][r])
				}
			}

			if j == -1 {
				dp[l][r] = 0
			}
		}
	}
	return dp[0][n-1] + 1
}
