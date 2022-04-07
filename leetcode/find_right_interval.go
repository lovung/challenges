package leetcode

import "sort"

type indexedInterval struct {
	index    int
	interval []int
}

func findRightInterval(intervals [][]int) []int {
	indexIntervals := make([]indexedInterval, len(intervals))
	for i, e := range intervals {
		indexIntervals[i] = indexedInterval{
			index:    i,
			interval: e,
		}
	}
	sort.SliceStable(indexIntervals, func(i, j int) bool {
		return indexIntervals[i].interval[0] < indexIntervals[j].interval[0]
	})
	result := make([]int, len(intervals))
loop_i:
	for i, e := range indexIntervals {
		for j := i; j < len(indexIntervals); j++ {
			if e.interval[1] > indexIntervals[j].interval[0] {
				continue
			}
			result[e.index] = indexIntervals[j].index
			continue loop_i
		}
		result[e.index] = -1
	}
	return result
}
