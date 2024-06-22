package jun2024

import (
	"slices"
	"sort"
)

// https://leetcode.com/problems/magnetic-force-between-two-balls/
func maxDistance(position []int, m int) int {
	sort.Ints(position)
	n := len(position)
	canArrange := func(d int) bool {
		val := position[0]
		cnt := 0
		for val <= position[n-1] {
			cnt++
			val += d
			idx, _ := slices.BinarySearch(position, val)
			if idx >= n {
				break
			}
			val = position[idx]
		}
		return cnt >= m
	}
	lo, hi := 1, (position[n-1]-position[0])/(m-1)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if canArrange(mid) {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi
}
