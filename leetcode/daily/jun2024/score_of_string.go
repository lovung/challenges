package jun2024

import "github.com/lovung/ds/maths"

// https://leetcode.com/problems/score-of-a-string/
func scoreOfString(s string) int {
	res := 0
	for i := 1; i < len(s); i++ {
		res += maths.ABS(int(s[i]) - int(s[i-1]))
	}
	return res
}
