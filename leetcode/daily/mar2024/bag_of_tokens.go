package mar2024

import (
	"sort"
)

// Link: https://leetcode.com/problems/bag-of-tokens/?envType=daily-question&envId=2024-03-04
func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	l, r := 0, len(tokens)-1
	score := 0
	for l <= r {
		if power >= tokens[l] {
			power -= tokens[l]
			score++
			l++
		} else if score > 0 && l < r && (power+tokens[r]) > tokens[l] {
			power += tokens[r]
			score--
			r--
		} else {
			break
		}
	}
	return score
}
