package biweekly130

// https://leetcode.com/problems/minimum-substring-partition-of-equal-character-frequency/description/
func minimumSubstringsInPartition(s string) int {
	sb := []byte(s)
	dp := make([]int, len(sb)+1)
	cnt := make([]int, 26)
	reset := [26]int{}
	for i := range sb {
		copy(cnt, reset[:])
		dp[i+1] = dp[i] + 1
		for j := i; j >= 0; j-- {
			cnt[sb[j]-'a']++
			if isBalanced(cnt) {
				dp[i+1] = min(dp[i+1], dp[j]+1)
			}
		}
	}
	return dp[len(sb)]
}

func isBalanced(cnt []int) bool {
	checkVal := 0
	for i := range cnt {
		if cnt[i] == 0 {
			continue
		}
		if checkVal > 0 && cnt[i] != checkVal {
			return false
		}
		checkVal = cnt[i]
	}
	return true
}
