package may2024

import "sort"

// https://leetcode.com/problems/boats-to-save-people/
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	cnt := 0
	for l, r := 0, len(people)-1; l <= r; r-- {
		if limit >= people[l]+people[r] {
			l++
		}
		cnt++
	}
	return cnt
}
