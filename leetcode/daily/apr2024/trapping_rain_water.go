package apr2024

// https://leetcode.com/problems/trapping-rain-water/
func trap(height []int) int {
	waterHeight := make([]int, len(height))
	curMax := 0
	for i := 0; i < len(height); i++ {
		curMax = max(curMax, height[i])
		waterHeight[i] = curMax
	}
	curMax = 0
	for i := len(height) - 1; i >= 0; i-- {
		curMax = max(curMax, height[i])
		waterHeight[i] = min(waterHeight[i], curMax)
	}
	sum := 0
	for i := range waterHeight {
		sum += waterHeight[i] - height[i]
	}
	return sum
}
