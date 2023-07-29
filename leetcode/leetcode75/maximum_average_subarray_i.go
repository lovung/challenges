package leetcode75

func findMaxAverage(nums []int, k int) float64 {
	slideSum := 0
	for i := 0; i < k; i++ {
		slideSum += nums[i]
	}
	maxSlideSum := slideSum
	for i := k; i < len(nums); i++ {
		slideSum += nums[i] - nums[i-k]
		if slideSum > maxSlideSum {
			maxSlideSum = slideSum
		}
	}
	return float64(maxSlideSum) / float64(k)
}
