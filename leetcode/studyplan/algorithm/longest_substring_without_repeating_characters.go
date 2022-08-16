package algorithm

func lengthOfLongestSubstring(s string) int {
	char := make([]int, 256)
	n := len(s)
	maxLen := 0
	for i := range char {
		char[i] = -1
	}
	for l, r := 0, 0; r < n; r++ {
		l = max(l, char[s[r]]+1)
		maxLen = max(maxLen, r-l+1)
		char[s[r]] = r
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
