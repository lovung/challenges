package problems

func maxUncrossedLines(A []int, B []int) int {
	var (
		nA, nB = len(A), len(B)
		dp     = make([][]int, 2)
		flag   = 1
	)
	dp[0] = make([]int, nB+1)
	dp[1] = make([]int, nB+1)

	for i := nA - 1; i >= 0; i-- {
		for j := nB - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[flag][j] = dp[flag^1][j+1] + 1
			} else {
				dp[flag][j] = max(dp[flag][j+1], dp[flag^1][j])
			}
		}
		flag ^= 1
	}
	return dp[flag^1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
