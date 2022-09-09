package medium

import "sort"

// Link: https://leetcode.com/problems/the-number-of-weak-characters-in-the-game/
func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] > properties[j][1]
		}
		return properties[i][0] < properties[j][0]
	})
	maxDefenseToEnd := make([]int, len(properties))
	max := 0
	for i := len(properties) - 1; i >= 0; i-- {
		if max < properties[i][1] {
			max = properties[i][1]
		}
		maxDefenseToEnd[i] = max
	}
	cnt := 0
	for i := range properties {
		if properties[i][1] < maxDefenseToEnd[i] {
			cnt++
		}
	}
	return cnt
}
