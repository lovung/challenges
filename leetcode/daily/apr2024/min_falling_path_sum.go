package apr2024

import "math"

// https://leetcode.com/problems/minimum-falling-path-sum-ii/
func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	dp := make([]int, n)
	for i := range grid {
		newdp := make([]int, n)
		prefixMin := calPrefixMin(dp)
		suffixMin := calSuffixMin(dp)
		for j := range grid[i] {
			minDpExceptJ := min(prefixMin[j], suffixMin[j])
			if prefixMin[j] == math.MaxInt && suffixMin[j] == math.MaxInt {
				minDpExceptJ = 0
			}
			newdp[j] = grid[i][j] + minDpExceptJ
		}
		copy(dp, newdp)
	}
	return minOf(dp, -1)
}

func calSuffixMin(row []int) []int {
	suffixMin := make([]int, len(row))
	minVal := math.MaxInt
	for i := len(row) - 1; i >= 0; i-- {
		suffixMin[i] = minVal
		minVal = min(minVal, row[i])
	}
	return suffixMin
}

func calPrefixMin(row []int) []int {
	prefixMin := make([]int, len(row))
	minVal := math.MaxInt
	for i := range row {
		prefixMin[i] = minVal
		minVal = min(minVal, row[i])
	}
	return prefixMin
}

func minOf(s []int, exceptIdx int) int {
	if len(s) == 1 {
		return s[0]
	}
	val := math.MaxInt
	for i := range s {
		if i != exceptIdx {
			val = min(val, s[i])
		}
	}
	return val
}
