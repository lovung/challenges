package problems

import "sort"

// Link: https://leetcode.com/problems/k-closest-points-to-origin/submissions/
func kClosest(points [][]int, K int) [][]int {
	var (
		n        = len(points)
		rangeArr = make([]int, n)
	)
	for i, e := range points {
		rangeArr[i] = e[0]*e[0] + e[1]*e[1]
	}
	sort.Ints(rangeArr)
	maxOkRange := rangeArr[K-1]
	var ans [][]int
	for _, e := range points {
		if e[0]*e[0]+e[1]*e[1] <= maxOkRange {
			ans = append(ans, e)
		}
	}
	return ans
}
