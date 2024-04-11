package apr2024

import (
	"slices"

	"golang.org/x/exp/constraints"
)

// https://leetcode.com/problems/minimum-area-rectangle/description/
func minAreaRect(points [][]int) int {
	xyMap := make(map[int][]int)
	yxMap := make(map[int][]int)
	for _, point := range points {
		xyMap[point[0]] = append(xyMap[point[0]], point[1])
		yxMap[point[1]] = append(yxMap[point[1]], point[0])
	}
	minArea := 0
	for xVal, yVals := range xyMap {
		for i := range yVals {
			for j := i + 1; j < len(yVals); j++ {
				y1, y2 := yVals[i], yVals[j]
				if len(yxMap[y1]) < 2 || len(yxMap[y2]) < 2 {
					continue
				}
				for _, k := range yxMap[y1] {
					if k <= xVal || !slices.Contains(yxMap[y2], k) {
						continue
					}
					if minArea == 0 {
						minArea = abs(k-xVal) * abs(y2-y1)
					} else {
						minArea = min(minArea, abs(k-xVal)*abs(y2-y1))
					}
				}
			}
		}
	}
	return minArea
}

func abs[T constraints.Float | constraints.Signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
