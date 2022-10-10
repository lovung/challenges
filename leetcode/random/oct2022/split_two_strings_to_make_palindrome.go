package oct2022

import "strings"

// Link: https://leetcode.com/problems/split-two-strings-to-make-palindrome/
func checkPalindromeFormation(a string, b string) bool {
	if isPalindrome(a) || isPalindrome(b) {
		return true
	}
	ra := reverseString(a)
	rb := reverseString(b)

	if checkPalindrome2NonPalindrome(a, b) ||
		checkPalindrome2NonPalindrome(b, a) {
		return true
	}
	if checkPalindrome2NonPalindrome(ra, rb) ||
		checkPalindrome2NonPalindrome(rb, ra) {
		return true
	}
	return false
}

// prefix_a + suffix_b
func checkPalindrome2NonPalindrome(a string, b string) bool {
	n := len(a)
	stopRight := false
	for l, r := 0, n-1; l < r && l < n/2; l++ {
		if !stopRight {
			if a[l] == b[r] {
				r--
				continue
			}
			stopRight = true
		}
		if stopRight && a[l] != a[n-1-l] {
			return false
		}
	}
	return true
}

func reverseString(s string) string {
	n := len(s)
	sb := strings.Builder{}
	for i := n - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l++ {
		if s[l] != s[r] {
			return false
		}
		l++
	}
	return true
}
