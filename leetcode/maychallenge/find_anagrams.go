package maychallenge

/*
 * Link: https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3332/
 */

func findAnagrams(s string, p string) []int {
	const a = 'a'
	lS := len(s)
	lP := len(p)
	pArr := make([]int, 26)
	sArr := make([]int, 26)
	var out []int

	for i := 0; i < lP; i++ {
		pArr[p[i]-a]++
	}
	for i := 0; i < lS; i++ {
		sArr[s[i]-a]++
		if i >= lP {
			sArr[s[i-lP]-a]--
		}
		if verify(sArr, pArr) {
			out = append(out, i-lP+1)
		}
	}
	return out
}

func verify(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
