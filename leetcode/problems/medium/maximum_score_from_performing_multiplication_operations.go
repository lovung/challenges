package medium

import (
	"github.com/lovung/ds/maths"
)

// Link: https://leetcode.com/problems/maximum-score-from-performing-multiplication-operations/
func maximumScore2(nums []int, multipliers []int) int {
	cache := make(map[cacheKey]int)
	return recursive(nums, multipliers, 0, len(nums)-1, cache)
}

type cacheKey struct {
	l, r      int
	remainLen int
}

func recursive(nums, multipliers []int, l, r int, cache map[cacheKey]int) int {
	if len(multipliers) == 1 {
		if multipliers[0] > 0 {
			return multipliers[0] * maths.Max(nums[l], nums[r])
		}
		return multipliers[0] * maths.Min(nums[l], nums[r])
	}
	if val, ok := cache[cacheKey{l, r, len(multipliers)}]; ok {
		return val
	}
	val := maths.Max(
		multipliers[0]*nums[l]+
			recursive(nums, multipliers[1:], l+1, r, cache),
		multipliers[0]*nums[len(nums)-1]+
			recursive(nums[:len(nums)-1], multipliers[1:], l, r-1, cache),
	)
	cache[cacheKey{l, r, len(multipliers)}] = val
	return val
}
