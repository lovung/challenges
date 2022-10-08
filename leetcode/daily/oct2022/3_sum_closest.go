package oct2022

import (
	"math"
	"sort"
)

// Link: https://leetcode.com/problems/3sum-closest/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := 0
	closest := math.MaxInt64
	n := len(nums)
	for first := 0; first < n-2; first++ {
		second, third := first+1, n-1
		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			diff := abs(sum - target)
			if diff < closest {
				closest = diff
				res = sum
			}
			if sum < target {
				second++
			} else {
				third--
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
