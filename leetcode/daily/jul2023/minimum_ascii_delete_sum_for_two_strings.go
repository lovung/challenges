package jul2023

import (
	"math"
)

// Link: https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings

// Solution 1: Timeout
func minimumDeleteSum(s1 string, s2 string) int {
	checkedSubString := make(map[string]bool)
	sum1 := sumASCII(s1)
	sum2 := sumASCII(s2)
	ret := math.MaxInt
	recursiveDelete(s1, s2, 0, sum1, sum2, 0, &ret, checkedSubString)
	return ret
}

func recursiveDelete(s, s2 string, i, sumS1, sumS2, curDeletedSum int, minRet *int, checkedSubString map[string]bool) {
	if checkedSubString[s] || checkIfCanBeTrimmedString(s, s2) {
		checkedSubString[s] = true
		*minRet = min(*minRet, curDeletedSum+sumS2-(sumS1-curDeletedSum))
	}
	if len(s) == 0 || i >= len(s) {
		return
	}

	// Delete char at index i
	recursiveDelete(s[:i]+s[i+1:], s2, i, sumS1, sumS2, curDeletedSum+int(s[i]), minRet, checkedSubString)
	// Don't delete char at index i
	recursiveDelete(s, s2, i+1, sumS1, sumS2, curDeletedSum, minRet, checkedSubString)
}

func sumASCII(s string) int {
	var sum int
	for i := range s {
		sum += int(s[i])
	}
	return sum
}

func checkIfCanBeTrimmedString(s string, s2 string) bool {
	if len(s) > len(s2) {
		return false
	}
	j := 0
	for i := range s {
		if j == len(s2) {
			return false
		}
		for ; j < len(s2); j++ {
			if s[i] != s2[j] {
				if j == len(s2)-1 {
					return false
				}
				continue
			}
			j++
			break
		}
	}
	return true
}

// Solution 2: Accepted
func minimumDeleteSum2(s1 string, s2 string) int {
	sum1 := sumASCII(s1)
	sum2 := sumASCII(s2)
	lcs := longestCommonSubsequenceMaxASCIISum(s1, s2)
	return sum1 + sum2 - 2*lcs
}

func longestCommonSubsequenceMaxASCIISum(s1 string, s2 string) int {
	dp := make([][]int, len(s1))
	for i := range dp {
		dp[i] = make([]int, len(s2))
	}
	for j := range s2 {
		if s1[0] == s2[j] {
			dp[0][j] = int(s1[0])
		} else if j > 0 {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := range s1 {
		if s2[0] == s1[i] {
			dp[i][0] = int(s2[0])
		} else if i > 0 {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i < len(s1); i++ {
		for j := 1; j < len(s2); j++ {
			if s1[i] != s2[j] {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			} else {
				dp[i][j] = dp[i-1][j-1] + int(s1[i])
			}
		}
	}
	return dp[len(s1)-1][len(s2)-1]
}
