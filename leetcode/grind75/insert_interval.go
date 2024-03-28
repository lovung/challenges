package grind75

import (
	"slices"
)

const (
	start = 0
	end   = 1
)

// Link: https://leetcode.com/problems/insert-interval/
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0, len(intervals)+1)
	mergedInteval := []int{newInterval[start], newInterval[end]}
	for i := range intervals {
		if intervals[i][start] > newInterval[end] && mergedInteval != nil {
			res = append(res, mergedInteval)
			res = append(res, intervals[i])
			mergedInteval = nil
			continue
		}
		if !overlap(intervals[i], newInterval) {
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
	return []int{min(a[start], b[start]), max(a[end], b[end])}
}

func overlap(a, b []int) bool {
	return !(a[end] < b[start] || a[start] > b[end])
}

// https://leetcode.com/problems/insert-interval/description/
func insert2(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	index, found := slices.BinarySearchFunc(intervals, newInterval[start], cmpFn)
	if !found {
		index = max(index-1, 0)
	}
	merged := tryMerge(intervals[index], newInterval)
	intervals[index] = merged[0]
	if len(merged) == 2 {
		index++
		// intervals = slices.Insert(intervals, index, merged[1])
		intervals = append(intervals, merged[1])
		copy(intervals[index+1:], intervals[index:])
		intervals[index] = merged[1]
	}

	// Refine remaining
	i := index + 1
	for ; i < len(intervals); i++ {
		merged := tryMerge(intervals[index], intervals[i])
		if len(merged) == 2 {
			break
		}
		intervals[index] = merged[0]
	}
	// intervals = slices.Delete(intervals, index+1, i)
	copy(intervals[index+1:], intervals[i:])
	return intervals[:len(intervals)-i+index+1]
}

func insert3(intervals [][]int, newInterval []int) [][]int {
	i, n := 0, len(intervals)
	for ; i < n && intervals[i][end] < newInterval[start]; i++ {
	}
	res := make([][]int, i, len(intervals)+1)
	copy(res, intervals[:i])
	for ; i < n && newInterval[end] >= intervals[i][start]; i++ {
		newInterval[start] = min(newInterval[start], intervals[i][start])
		newInterval[end] = max(newInterval[end], intervals[i][end])
	}
	res = append(res, newInterval)
	newRes := make([][]int, len(res)+len(intervals[i:]))
	copy(newRes[:], res)
	copy(newRes[len(res):], intervals[i:])
	return newRes
}

func cmpFn(s []int, t int) int {
	if s[0] < t {
		return -1
	}
	if s[0] > 1 {
		return 1
	}
	return 0
}

func tryMerge(a, b []int) [][]int {
	if a[start] == b[start] {
		return [][]int{{a[start], max(a[end], b[end])}}
	}
	if a[start] > b[start] {
		a, b = b, a
	}
	if a[end] < b[start] {
		return [][]int{a, b}
	}
	return [][]int{{a[start], max(a[end], b[end])}}
}
