package jul2023

// https://leetcode.com/problems/longest-common-subsequence
func longestCommonSubsequence(s1 string, s2 string) int {
	dp := make([][]int, len(s1))
	for i := range dp {
		dp[i] = make([]int, len(s2))
	}
	for j := range s2 {
		if s1[0] == s2[j] {
			dp[0][j] = 1
		} else if j > 0 {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := range s1 {
		if s2[0] == s1[i] {
			dp[i][0] = 1
		} else if i > 0 {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i < len(s1); i++ {
		for j := 1; j < len(s2); j++ {
			if s1[i] != s2[j] {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			} else {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}
	return dp[len(s1)-1][len(s2)-1]
}
