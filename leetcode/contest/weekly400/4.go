package weekly400

import (
	"math"

	"github.com/lovung/ds/maths"
)

func minimumDifference(nums []int, k int) int {
	res := math.MaxInt
	dp := [100001]int{}
	for i := range nums {
		and := nums[i]
		for j := i; j < len(nums); j++ {
			and &= nums[j]
			res = min(res, maths.ABS(k-and))
			if k-and >= res || and <= dp[j] {
				break
			}
			dp[j] = and
		}
	}
	return res
}
