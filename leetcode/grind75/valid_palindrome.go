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

func isUpperChar(b byte) bool {
	return 'A' <= b && 'Z' >= b
}

func isLowerOrNumChar(b byte) bool {
	return ('a' <= b && 'z' >= b) ||
		('0' <= b && '9' >= b)
}

func isPalindrome2(s string) bool {
	var lc, rc byte
	for l, r := 0, len(s)-1; l < r; {
		lc = s[l]
		if isUpperChar(s[l]) {
			lc = lc - 'A' + 'a'
		}
		if !isLowerOrNumChar(lc) {
			l++
			continue
		}
		rc = s[r]
		if isUpperChar(rc) {
			rc = rc - 'A' + 'a'
		}
		if !isLowerOrNumChar(rc) {
			r--
			continue
		}
		if lc != rc {
			return false
		}
		l++
		r--
	}
	return true
}
