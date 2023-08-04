package dp

import (
	"github.com/lovung/ds/matrix"
)

// https://leetcode.com/problems/unique-paths-ii/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := matrix.New[int](len(obstacleGrid), len(obstacleGrid[0]))
	dp.UpdateAt(0, 0, 1)
	for i := range obstacleGrid {
		for j := range obstacleGrid[i] {
			if obstacleGrid[i][j] == 1 {
				dp.UpdateAt(i, j, 0)
				continue
			}
			if i == 0 && j == 0 {
				continue
			}
			dp.UpdateAt(i, j, dp.At(i-1, j)+dp.At(i, j-1))
		}
	}
	return dp.At(len(dp.Data())-1, len(dp.Data()[0])-1)
}
