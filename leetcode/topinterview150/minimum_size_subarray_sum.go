package topinterview150

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
		for j := min(len(presum)-1, i+minSubArrLen); j > i; j-- {
			if presum[j]-presum[i] < target {
				break
			}
			minSubArrLen = min(minSubArrLen, j-i)
		}
	}
	return minSubArrLen
}
