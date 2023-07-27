package grind75

// https://leetcode.com/problems/valid-anagram/
func isAnagram(s string, t string) bool {
	charCnt := make([]int, 26)
	for i := range s {
		charCnt[s[i]-'a']++
	}
	for i := range t {
		charCnt[t[i]-'a']--
	}
	for i := range charCnt {
		if charCnt[i] != 0 {
			return false
		}
	}
	return true
}
