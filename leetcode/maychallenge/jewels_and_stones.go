package maychallenge

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
