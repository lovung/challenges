package problems

func minDistance(word1 string, word2 string) int {
	var (
		len1 = len(word1)
		len2 = len(word2)
		dp   = make([][]int, 2)
	)
	dp[0] = make([]int, len1+1)
	dp[1] = make([]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= len2; i++ {
		for j := 0; j <= len1; j++ {
			if j == 0 {
				dp[i%2][j] = i
			} else if word1[j-1] == word2[i-1] {
				dp[i%2][j] = dp[(i-1)%2][j-1]
			} else {
				dp[i%2][j] = 1 + min(dp[(i-1)%2][j], min(dp[i%2][j-1], dp[(i-1)%2][j-1]))
			}
		}
	}
	return dp[len2%2][len1]
}
