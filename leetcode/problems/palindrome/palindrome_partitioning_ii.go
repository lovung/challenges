package palindrome

// Link: https://leetcode.com/problems/palindrome-partitioning-ii/
// very slow solution
// pass: 24/36
func minCut(s string) int {
	res := 2000
	dfsCountMinCut(&res, 0, s)
	return res
}

func dfsCountMinCut(minCut *int, cutCnt int, remainStr string) {
	if len(remainStr) == 0 {
		if cutCnt-1 < *minCut {
			*minCut = cutCnt - 1
		}
		return
	}
	for end := 1; end <= len(remainStr); end++ {
		if isPalindrome(remainStr[:end]) {
			dfsCountMinCut(minCut, cutCnt+1, remainStr[end:])
		}
	}
}

// DP solution
// Time: O(N^3)
// Space: O(N^2)
// Pass: 32/36
func minCut2(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	cntK := 0
	for startJ := 1; startJ < n; startJ++ {
		for i := 0; i < n-startJ; i++ {
			j := i + startJ
			if dp[i+1][j-1] == 0 && s[i] == s[j] {
				continue
			}
			if dp[i][j-1] == 0 || dp[i+1][j] == 0 {
				dp[i][j]++
				continue
			}
			if dp[i+1][j-1] == 0 && s[i] != s[j] {
				dp[i][j] = 2
				continue
			}
			cntK += j - i
			min := dp[i][i] + dp[i+1][j]
			for k := 1; k < j-i; k++ {
				tmp := dp[i][i+k] + dp[i+1+k][j]
				if min > tmp {
					min = tmp
				}
				if min == 0 {
					break
				}
			}
			dp[i][j] = min + 1
		}
	}
	return dp[0][n-1]
}

func minCut3(s string) int {
	n := len(s)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = i
	}
	for mid := 1; mid < n; mid++ {
		for start, end := mid, mid; start >= 0 && end < n && s[start] == s[end]; start, end = start-1, end+1 {
			newCutAtEnd := 0
			if start != 0 {
				newCutAtEnd = dp[start-1] + 1
			}
			dp[end] = min(dp[end], newCutAtEnd)
		}
		for start, end := mid-1, mid; start >= 0 && end < n && s[start] == s[end]; start, end = start-1, end+1 {
			newCutAtEnd := 0
			if start != 0 {
				newCutAtEnd = dp[start-1] + 1
			}
			dp[end] = min(dp[end], newCutAtEnd)
		}
	}
	return dp[n-1]
}
