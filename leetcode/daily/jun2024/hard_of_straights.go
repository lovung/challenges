package jun2024

import (
	"slices"
	"sort"
)

// https://leetcode.com/problems/hand-of-straights
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	sort.Ints(hand)
	groups := make(map[int][]int)
	groupCnt := 0
outter:
	for _, h := range hand {
		for i, size := range groups[h-1] {
			if size < groupSize {
				groups[h-1] = slices.Delete(groups[h-1], i, i+1)
				groups[h] = append(groups[h], size+1)
				continue outter
			}
		}
		if groupCnt >= len(hand)/groupSize {
			return false
		}
		groups[h] = append(groups[h], 1)
		groupCnt++
	}
	return true
}
