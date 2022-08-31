package problems

import "sort"

type indexedInterval struct {
	index    int
	interval []int
}

func findRightInterval(intervals [][]int) []int {
	indexIntervals := make([]*indexedInterval, len(intervals))
	for i, e := range intervals {
		indexIntervals[i] = &indexedInterval{
			index:    i,
			interval: e,
		}
	}
	sort.SliceStable(indexIntervals, func(i, j int) bool {
		return indexIntervals[i].interval[0] < indexIntervals[j].interval[0]
	})
	result := make([]int, len(intervals))
	for i, e := range indexIntervals {
		val := e.interval[1]
		idx := i
		additionalIndex := binarySearchForIndexedInterval(indexIntervals[i:], func(j int) int {
			if val < indexIntervals[idx+j].interval[0] {
				return -1
			} else if val > indexIntervals[idx+j].interval[0] {
				return 1
			}
			return 0
		})
		if additionalIndex != -1 {
			result[e.index] = indexIntervals[idx+additionalIndex].index
		} else {
			result[e.index] = -1
		}
	}
	return result
}

// make sure the slice is sorted
// compFunc returns 0, -1 or 1
func binarySearchForIndexedInterval(slice []*indexedInterval, compFunc func(i int) int) int {
	l, r := 0, len(slice)-1
	pivot := 0
	for l <= r {
		pivot = (l + r) / 2
		cmpVal := compFunc(pivot)
		switch cmpVal {
		case -1:
			r = pivot - 1
		case 1:
			l = pivot + 1
		case 0:
			return pivot
		}
	}
	if l >= len(slice) {
		return -1
	}
	return l
}
