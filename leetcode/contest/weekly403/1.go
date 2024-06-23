package weekly403

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/minimum-average-of-smallest-and-largest-elements/description/
func minimumAverage(nums []int) float64 {
	sort.Ints(nums)
	res := math.MaxFloat64
	for i := range len(nums) / 2 {
		res = min(res, float64(nums[i]+nums[len(nums)-1-i])/2)
	}
	return res
}
