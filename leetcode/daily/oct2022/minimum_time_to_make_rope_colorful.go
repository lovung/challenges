package oct2022

import "github.com/lovung/ds/maths"

// Link: https://leetcode.com/problems/minimum-time-to-make-rope-colorful/
func minCost(colors string, neededTime []int) int {
	res := 0
	totalCostSameColor := 0
	maxCostSameColor := 0
	prevColor := byte(0)
	for i := range colors {
		if colors[i] != prevColor {
			// change color
			res += totalCostSameColor - maxCostSameColor
			totalCostSameColor = 0
			maxCostSameColor = 0
		}
		totalCostSameColor += neededTime[i]
		maxCostSameColor = maths.Max(maxCostSameColor, neededTime[i])
		prevColor = colors[i]
	}
	res += totalCostSameColor - maxCostSameColor
	return res
}
