package contest394

import "math"

func minimumOperations(grid [][]int) int {
	n := len(grid)
	prevDp := make([]int, 10)
	dp := make([]int, 10)
	for j := range grid[0] {
		for k := range dp {
			dp[k] = n
		}
		for i := range grid {
			dp[grid[i][j]]--
		}
		for k := range dp {
			dp[k] += minOf(prevDp, k)
		}
		copy(prevDp, dp)
	}
	return minOf(dp, -1)
}

func minOf(s []int, exceptIdx int) int {
	val := math.MaxInt
	for i := range s {
		if i != exceptIdx {
			val = min(val, s[i])
		}
	}
	return val
}
