package biweekly130

import "github.com/lovung/ds/matrix"

// https://leetcode.com/problems/check-if-grid-satisfies-conditions/description/
func satisfiesConditions(grid [][]int) bool {
	for i := range grid {
		for j := range grid[i] {
			if matrix.In(grid, i+1, j) && grid[i][j] != grid[i+1][j] {
				return false
			}
			if matrix.In(grid, i, j+1) && grid[i][j] == grid[i][j+1] {
				return false
			}
		}
	}
	return true
}
