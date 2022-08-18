package algorithm

var (
	isBadVersion = func(k int) bool { return k >= 4 }
)

// Link: https://leetcode.com/problems/first-bad-version/
// BigO: O(log n)
func firstBadVersion(n int) int {
	l, r := 0, n
	for {
		m := (l + r) / 2
		tmp := isBadVersion(m)
		if !tmp && isBadVersion(m+1) {
			return m + 1
		}
		if tmp {
			r = m
		} else {
			l = m
		}
	}
}
