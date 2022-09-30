package grind75

import (
	"sort"

	"github.com/lovung/ds/maths"
)

// Link: https://leetcode.com/problems/merge-intervals/
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	if len(intervals) == 0 {
		return res
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else {
			res[len(res)-1] = []int{
				res[len(res)-1][0],
				maths.Max(intervals[i][1], res[len(res)-1][1]),
			}
		}
	}
	return res
}
