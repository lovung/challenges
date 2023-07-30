package leetcode75

import "strings"

// https://leetcode.com/problems/greatest-common-divisor-of-strings
func gcdOfStrings(str1 string, str2 string) string {
	if str1 == str2 {
		return str1
	}
	if len(str1) > len(str2) {
		str1, str2 = str2, str1
	}
	if !strings.HasPrefix(str2, str1) {
		return ""
	}
	return gcdOfStrings(strings.TrimPrefix(str2, str1), str1)
}
