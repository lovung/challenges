package oct2022

import "strconv"

// Link: https://leetcode.com/problems/decode-ways/
func numDecodings(s string) int {
	return numDecodingsBytes([]byte(s))
}

// O(2^N)
func numDecodingsBytes(s []byte) int {
	if len(s) == 0 {
		return 1
	}
	if len(s) == 1 {
		if s[0] == '0' {
			return 0
		} else {
			return 1
		}
	}
	switch s[0] {
	case '0':
		return 0
	case '1':
		return numDecodingsBytes(s[1:]) + numDecodingsBytes(s[2:])
	case '2':
		res := numDecodingsBytes(s[1:])
		switch s[1] {
		case '0', '1', '2', '3', '4', '5', '6':
			res += numDecodingsBytes(s[2:])
		}
		return res
	default:
		return numDecodingsBytes(s[1:])
	}
}

// O(N)
func numDecodings2(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	}
	for i := 2; i <= n; i++ {
		first, _ := strconv.Atoi(s[i-1 : i])
		second, _ := strconv.Atoi(s[i-2 : i])
		if first >= 1 && first <= 9 {
			dp[i] += dp[i-1]
		}
		if second >= 10 && second <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
