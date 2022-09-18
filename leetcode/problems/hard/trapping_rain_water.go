package hard

import "github.com/lovung/ds/maths"

// Link: https://leetcode.com/problems/trapping-rain-water/
func trap(height []int) int {
	n := len(height)
	highestLtoR := make([]int, n)
	highestRtoL := make([]int, n)

	highest := height[0]
	for i := range height {
		if height[i] > highest {
			highest = height[i]
		}
		highestLtoR[i] = highest
	}
	highest = height[n-1]
	for i := n - 1; i >= 0; i-- {
		if height[i] > highest {
			highest = height[i]
		}
		highestRtoL[i] = highest
	}

	trappedWater := 0
	for i := range height {
		trappedWater += maths.Min(highestLtoR[i], highestRtoL[i]) - height[i]
	}
	return trappedWater
}
