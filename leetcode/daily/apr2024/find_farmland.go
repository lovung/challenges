package apr2024

import "github.com/lovung/ds/matrix"

// https://leetcode.com/problems/find-all-groups-of-farmland/
func findFarmland(land [][]int) [][]int {
	topLeftMatrix := matrix.New[point](len(land), len(land[0]))
	res := make([][]int, 0)
	for i := range land {
		for j := range land[i] {
			if land[i][j] == 0 {
				continue
			}
			if isTopLeft(land, i, j) {
				topLeftMatrix.UpdateAt(i, j, point{i, j})
			} else if matrix.In(land, i-1, j) && land[i-1][j] == 1 {
				topLeftMatrix.UpdateAt(i, j, topLeftMatrix.At(i-1, j))
			} else if matrix.In(land, i, j-1) && land[i][j-1] == 1 {
				topLeftMatrix.UpdateAt(i, j, topLeftMatrix.At(i, j-1))
			}
			if isBottomRight(land, i, j) {
				farmland := make([]int, 0, 4)
				tl := topLeftMatrix.At(i, j)
				farmland = append(farmland, tl.x, tl.y, i, j)
				res = append(res, farmland)
			}
		}
	}
	return res
}

func isTopLeft(land [][]int, i, j int) bool {
	return (!matrix.In(land, i-1, j) || land[i-1][j] == 0) &&
		(!matrix.In(land, i, j-1) || land[i][j-1] == 0)
}

func isBottomRight(land [][]int, i, j int) bool {
	return (!matrix.In(land, i+1, j) || land[i+1][j] == 0) &&
		(!matrix.In(land, i, j+1) || land[i][j+1] == 0)
}
