package junechallenge

// Link: https://leetcode.com/explore/featured/card/june-leetcoding-challenge/540/week-2-june-8th-june-14th/3355/
func isSubsequence(s string, t string) bool {
	var (
		r1, r2 int
		n1, n2 = len(s), len(t)
	)
	for r1 < n1 && r2 < n2 {
		if s[r1] == t[r2] {
			r1++
		}
		r2++
	}
	return r1 == n1
}
