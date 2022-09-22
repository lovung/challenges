package grind75

import "github.com/lovung/ds/maths"

// Link: https://leetcode.com/problems/insert-interval/
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0, len(intervals)+1)
	mergedInteval := []int{newInterval[0], newInterval[1]}
	for i := range intervals {
		if intervals[i][0] > newInterval[1] && mergedInteval != nil {
			res = append(res, mergedInteval)
			res = append(res, intervals[i])
			mergedInteval = nil
			continue
		}
		if intervals[i][1] < newInterval[0] || intervals[i][0] > newInterval[1] {
			res = append(res, intervals[i])
			continue
		}
		mergedInteval = mergeIntervals(mergedInteval, intervals[i])
	}
	if mergedInteval != nil {
		res = append(res, mergedInteval)
	}
	return res
}

func mergeIntervals(a, b []int) []int {
	return []int{maths.Min(a[0], b[0]), maths.Max(a[1], b[1])}
}
