package grind75

import "strings"

// Link: https://leetcode.com/problems/valid-palindrome/
func isPalindrome(s string) bool {
	return isPalindromePure(removeNonAlphanumeric(s))
}

func removeNonAlphanumeric(s string) string {
	sb := strings.Builder{}
	for i := range s {
		if ('a' <= s[i] && 'z' >= s[i]) ||
			('0' <= s[i] && '9' >= s[i]) {
			sb.WriteByte(s[i])
		}
		if 'A' <= s[i] && 'Z' >= s[i] {
			sb.WriteByte(s[i] - 'A' + 'a')
		}
	}
	return sb.String()
}

func isPalindromePure(s string) bool {
	for l, r := 0, len(s)-1; l < r; {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
