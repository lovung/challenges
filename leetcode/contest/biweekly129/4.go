package biweekly129

type pair struct {
	zero, one int
}

func (p *pair) sum() int {
	return p.zero + p.one
}

const mod = 1e9 + 7

func numberOfStableArrays2(zero int, one int, limit int) int {
	dp := make([][]pair, zero+1)
	for i := range dp {
		dp[i] = make([]pair, one+1)
	}
	for i := 0; i <= zero; i++ {
		if i > limit {
			break
		}
		dp[i][0] = pair{1, 0}
	}
	for j := 0; j <= one; j++ {
		if j > limit {
			break
		}
		dp[0][j] = pair{0, 1}
	}
	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			dp[i][j].zero = dp[i-1][j].sum()
			dp[i][j].one = dp[i][j-1].sum()
			if i > limit {
				dp[i][j].zero -= dp[i-limit-1][j].one
			}
			if j > limit {
				dp[i][j].one -= dp[i][j-limit-1].zero
			}
			dp[i][j].zero = (dp[i][j].zero + mod) % mod
			dp[i][j].one = (dp[i][j].one + mod) % mod
		}
	}
	return dp[zero][one].sum() % mod
}
