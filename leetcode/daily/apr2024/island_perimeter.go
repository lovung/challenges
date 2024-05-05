package apr2024

import "github.com/lovung/ds/matrix"

// https://leetcode.com/problems/island-perimeter/
func islandPerimeter(grid [][]int) int {
	perimeter := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			for _, d := range matrix.Dirs {
				x := i + d[0]
				y := j + d[1]
				if !matrix.In(grid, x, y) || grid[x][y] == 0 {
					perimeter++
				}
			}
		}
	}
	return perimeter
}
