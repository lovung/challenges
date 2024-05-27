package biweekly130

import (
	"slices"

	"github.com/lovung/ds/maths"
)

type point struct {
	x, y  int
	tag   byte
	level int
}

func maxPointsInsideSquare(points [][]int, s string) int {
	pointList := make([]point, 0, len(points))
	for i, p := range points {
		pointList = append(pointList, point{p[0], p[1], s[i], max(maths.ABS(p[0]), maths.ABS(p[1]))})
	}
	slices.SortFunc(pointList, func(a, b point) int { return a.level - b.level })
	res, cnt := 0, 0
	exist := make(map[byte]bool)
	for i, p := range pointList {
		if i > 0 && p.level > pointList[i-1].level {
			res = cnt
		}
		if exist[p.tag] {
			return res
		}
		exist[p.tag] = true
		cnt++
	}
	return cnt
}
