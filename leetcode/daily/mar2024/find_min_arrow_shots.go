package mar2024

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	arrEnd, cnt := math.MinInt, 0
	for i := range points {
		if points[i][0] > arrEnd {
			cnt++
			arrEnd = points[i][1]
		}
	}
	return cnt
}
