package contest

import (
	"math"

	"github.com/lovung/ds/maths"
)

// https://leetcode.com/problems/sorting-three-groups/
func minimumOperations(nums []int) int {
	min := math.MaxInt
	for i := 0; i <= len(nums); i++ {
		for j := i; j <= len(nums); j++ {
			group1 := countMatch(nums[:i], 1)
			group2 := countMatch(nums[i:j], 2)
			group3 := countMatch(nums[j:], 3)
			min = maths.Min(min, group1+group2+group3)
		}
	}
	return min
}

func countMatch(nums []int, val int) int {
	cnt := 0
	for i := range nums {
		if nums[i] != val {
			cnt++
		}
	}
	return cnt
}
