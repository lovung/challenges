package may2024

const mod = 1e9 + 7

// https://leetcode.com/problems/student-attendance-record-ii/
// Wrong solution... still don't know why
func checkRecord_wrong(n int) int {
	type cases struct{ all, absent int }
	dp := make([][3]cases, n+1)
	dp[1] = [3]cases{{1, 1}, {1, 0}, {1, 0}}
	for i := 2; i <= n; i++ {
		prevAbsent := (dp[i-1][0].absent + dp[i-1][1].absent + dp[i-1][2].absent) % mod
		prevTotal := (dp[i-1][0].all + dp[i-1][1].all + dp[i-1][2].all) % mod
		dp[i][0].all = (prevTotal - prevAbsent) % mod
		dp[i][0].absent = (prevTotal - prevAbsent) % mod
		dp[i][1].all = (prevTotal - dp[i-2][1].all)
		dp[i][1].absent = (prevAbsent - dp[i-2][1].absent)
		if i > 2 {
			dp[i][1].all += dp[i-3][1].all
			dp[i][1].all %= mod
			dp[i][1].absent += dp[i-3][1].absent
			dp[i][1].absent %= mod
		}
		dp[i][2].all = prevTotal % mod
		dp[i][2].absent = prevAbsent % mod
	}
	return (dp[n][0].all + dp[n][1].all + dp[n][2].all)
}

func checkRecord(n int) int {
	// Initialize dp array
	dp := [2][3]int{}
	dp[0][0] = 1

	for i := 0; i < n; i++ {
		prev := dp // Make a copy of the current state of d
		// Update dp for the next state
		// add P only
		dp[0][0] = (prev[0][0] + prev[0][1] + prev[0][2]) % mod
		// add L
		dp[0][1] = prev[0][0]
		// add L
		dp[0][2] = prev[0][1]
		// add A or P
		dp[1][0] = (prev[0][0] + prev[0][1] + prev[0][2] +
			prev[1][0] + prev[1][1] + prev[1][2]) % mod
		// add L
		dp[1][1] = prev[1][0]
		// add L
		dp[1][2] = prev[1][1]
	}

	// Calculate the final result
	result := (dp[0][0] + dp[0][1] + dp[0][2] + dp[1][0] + dp[1][1] + dp[1][2]) % mod
	return result
}
