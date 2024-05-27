package weekly397

import "github.com/lovung/ds/maths"

// https://leetcode.com/problems/permutation-difference-between-two-strings/description/
func findPermutationDifference(s string, t string) int {
	const a = 'a'
	idx := [26]int{}
	for i := range s {
		idx[s[i]-a] = i
	}
	res := 0
	for i := range t {
		res += maths.ABS(idx[t[i]-a] - i)
	}
	return res
}
