package problems

/*
 * Link: https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3317/
 */

func numJewelsInStones(J string, S string) int {
	// Map solution
	count := 0
	stones := make(map[rune]int)
	for _, e := range S {
		stones[e]++
	}
	for _, e := range J {
		count += stones[e]
	}
	return count
}
