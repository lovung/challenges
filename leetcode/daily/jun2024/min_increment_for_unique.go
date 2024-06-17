package jun2024

import (
	"math"
)

// https://leetcode.com/problems/minimum-increment-to-make-array-unique
func minIncrementForUnique(nums []int) int {
	minVal, maxVal := math.MaxInt, 0
	cnt := make([]int, 2*1e5)
	for i := range nums {
		cnt[nums[i]]++
		maxVal = max(maxVal, nums[i])
		minVal = min(minVal, nums[i])
	}
	r := minVal
	for cnt[r] > 0 {
		r++
	}
	res := 0
	for l := minVal; l <= maxVal; l++ {
		if cnt[l] <= 1 {
			continue
		}
		if r < l {
			for r = l + 1; cnt[r] > 0; r++ {
			}
		}
		for range cnt[l] - 1 {
			res += r - l
			for r = r + 1; cnt[r] > 0; r++ {
			}
		}
	}
	return res
}
