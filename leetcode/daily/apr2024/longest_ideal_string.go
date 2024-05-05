package apr2024

// https://leetcode.com/problems/longest-ideal-subsequence
func longestIdealString(s string, k int) int {
	const a byte = 'a'
	dp := make([]int, 26)
	for i := range s {
		maxVal := 0
		idx := int(s[i] - a)
		for j := max(0, idx-k); j <= min(idx+k, 25); j++ {
			maxVal = max(maxVal, dp[j])
		}
		dp[idx] = maxVal + 1
	}
	res := 0
	for i := range dp {
		res = max(res, dp[i])
	}
	return res
}
