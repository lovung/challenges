package leetcode75

// https://leetcode.com/problems/merge-strings-alternately
func mergeAlternately(word1 string, word2 string) string {
	isSwap := false
	if len(word1) > len(word2) {
		word1, word2 = word2, word1
		isSwap = true
	}
	merged := ""
	i := 0
	for ; i < len(word1); i++ {
		if isSwap {
			merged += string(word2[i])
			merged += string(word1[i])
		} else {
			merged += string(word1[i])
			merged += string(word2[i])
		}
	}
	return merged + word2[i:]
}
