package may2024

import (
	"math"

	"github.com/lovung/ds/maths"
)

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	total, totalDiff := int64(0), int64(0)
	for i := range nums {
		total += int64(nums[i])
	}
	positiveCnt := 0
	minAbsDiff := math.MaxInt
	for _, n := range nums {
		diff := (n ^ k) - n
		if diff > 0 {
			totalDiff += int64(diff)
			positiveCnt++
		}
		minAbsDiff = min(minAbsDiff, maths.ABS(diff))
	}
	if positiveCnt%2 == 1 {
		totalDiff -= int64(minAbsDiff)
	}
	return total + totalDiff
}
