package medium

import "sort"

// Link: https://leetcode.com/problems/bag-of-tokens/
func bagOfTokensScore(tokens []int, power int) int {
	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i] < tokens[j]
	})
	l, r := 0, len(tokens)-1
	score := 0
	for l <= r {
		if power >= tokens[l] {
			power -= tokens[l]
			score++
			l++
		} else if score > 0 && l < r && (power + tokens[r]) > tokens[l] {
			power += tokens[r]
			score--
			r--
		} else {
			break
		}
	}
	return score
}
