package algorithm

import "strings"

// Link: https://leetcode.com/problems/reverse-string/
// BigO: O(n)
func reverseString(s []byte) {
	n2 := len(s) / 2
	for i := 0; i < n2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

// Link: https://leetcode.com/problems/reverse-words-in-a-string/
// BigO: O(n)
func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	for i := range ss {
		b := []byte(ss[i])
		reverseString(b)
		ss[i] = string(b)
	}
	return strings.Join(ss, " ")
}
