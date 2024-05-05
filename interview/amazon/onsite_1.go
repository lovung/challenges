package amazon

import "github.com/lovung/ds/matrix"

func CountGroupOfHouseType(grid [][]string, target string) int {
	counter := 0
	for i := range grid {
		once := false
		for j := range grid[i] {
			if grid[i][j] != target {
				continue
			}
			if !matrix.In(grid, i-1, j) || grid[i-1][j] != target {
				// First target in the row
				counter++
				once = true
			}
			if matrix.In(grid, i, j-1) && grid[i][j-1] == target {
				if once == true {
					once = false
					counter--
				}
			}
		}
	}
	return counter
}
