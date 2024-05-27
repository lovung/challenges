package may2024

import "github.com/lovung/ds/matrix"

// https://leetcode.com/problems/path-with-maximum-gold
func getMaximumGold(grid [][]int) int {
	type pos struct{ x, y int }
	roots := make([]pos, 0, 25)
	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			switch countAdj(grid, i, j) {
			case 0:
				res = max(res, grid[i][j])
			case 1, 2:
				roots = append(roots, pos{i, j})
			}
		}
	}
	for _, l := range roots {
		visited := matrix.New[bool](len(grid), len(grid[0]))
		dfs(grid, visited, l.x, l.y, &res, 0)
	}
	return res
}

func dfs(grid [][]int, visited *matrix.Matrix[bool], i, j int, maxCollected *int, collected int) {
	collected += grid[i][j]
	*maxCollected = max(*maxCollected, collected)
	visited.UpdateAt(i, j, true)
	for _, d := range matrix.Dirs {
		x, y := i+d[0], j+d[1]
		if !matrix.In(grid, x, y) || grid[x][y] == 0 || visited.At(x, y) {
			continue
		}
		dfs(grid, visited, x, y, maxCollected, collected)
	}
	visited.UpdateAt(i, j, false)
}

func countAdj(grid [][]int, i, j int) int {
	cnt := 0
	for _, d := range matrix.Dirs {
		x, y := i+d[0], j+d[1]
		if matrix.In(grid, x, y) && grid[x][y] > 0 {
			cnt++
		}
	}
	return cnt
}
