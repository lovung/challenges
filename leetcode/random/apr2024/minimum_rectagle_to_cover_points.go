package apr2024

import "sort"

// https://leetcode.com/problems/minimum-rectangles-to-cover-points/
func minRectanglesToCoverPoints(points [][]int, w int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	cnt, lastX := 0, -1
	for _, p := range points {
		if lastX < 0 || lastX+w < p[0] {
			cnt++
			lastX = p[0]
		}
	}
	return cnt
}
