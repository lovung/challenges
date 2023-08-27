package contest

import (
	"math"

	"github.com/lovung/ds/maths"
)

// https://leetcode.com/problems/minimum-absolute-difference-between-elements-with-constraint/
func minAbsoluteDifference(nums []int, x int) int {
	n := len(nums)
	ret := math.MaxInt
	for i := 0; i < n-x; i++ {
		for j := i + x; j < n; j++ {
			ret = maths.Min(ret, abs(nums[i]-nums[j]))
			if ret == 0 {
				return ret
			}
		}
	}
	return ret
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
