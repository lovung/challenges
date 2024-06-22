package jun2024

import "slices"

// https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets
func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	canBloom := func(bDay int) bool {
		bouquet, cnt := 0, 0
		for _, b := range bloomDay {
			if b <= bDay {
				cnt++
			} else {
				cnt = 0
			}
			if cnt == k {
				cnt = 0
				bouquet++
			}
		}
		return bouquet >= m
	}
	maxBloomDay := slices.Max(bloomDay)
	minBloomDay := slices.Min(bloomDay)
	lo, hi := minBloomDay, maxBloomDay
	for lo <= hi {
		mid := (lo + hi) >> 1 // /2
		if canBloom(mid) {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}
