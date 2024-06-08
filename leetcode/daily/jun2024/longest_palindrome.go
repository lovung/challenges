package jun2024

// https://leetcode.com/problems/longest-palindrome
func longestPalindrome(s string) int {
	cnt := [52]int{}
	for i := range s {
		if 'a' <= s[i] && s[i] <= 'z' {
			cnt[s[i]-'a'+26]++
		} else {
			cnt[s[i]-'A']++
		}
	}
	res := 0
	oddtaken := false
	for i := range cnt {
		if cnt[i]&1 == 0 {
			res += cnt[i]
		} else if !oddtaken {
			res += cnt[i]
			oddtaken = true
		} else {
			res += cnt[i] - 1
		}
	}
	return res
}
