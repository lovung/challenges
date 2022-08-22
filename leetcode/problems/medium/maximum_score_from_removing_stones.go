package medium

// Link: https://leetcode.com/problems/maximum-score-from-removing-stones/
func maximumScore(a int, b int, c int) int {
	// bubble sort
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}

	// triagle rule check
	if a+b < c {
		return a + b
	}
	return (a + b + c) / 2
}
