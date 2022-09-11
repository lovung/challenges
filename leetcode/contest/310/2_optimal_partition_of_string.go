package contest

// Link: https://leetcode.com/problems/optimal-partition-of-string/
func partitionString(s string) int {
	subCnt := 1
	cnt := make([]int, 26)
	for i := range s {
		if cnt[s[i]-'a'] > 0 {
			subCnt++
			cnt = make([]int, 26)
		}
		cnt[s[i]-'a']++
	}
	return subCnt
}
