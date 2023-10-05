package topinterview150

import "github.com/lovung/ds/maths"

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	presum := make([]int, len(nums)+1)
	for i := range nums {
		sum += nums[i]
		presum[i+1] = sum
	}
	if sum < target {
		return 0
	}
	minSubArrLen := len(nums)
	for i := 0; i < len(presum)-1; i++ {
		for j := maths.Min(len(presum) - 1, i + minSubArrLen); j > i; j-- {
			if presum[j]-presum[i] < target {
				break
			}
			minSubArrLen = maths.Min(minSubArrLen, j-i)
		}
	}
	return minSubArrLen
}
