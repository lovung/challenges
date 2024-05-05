package apr2024

// https://leetcode.com/problems/magic-squares-in-grid/description/
func numMagicSquaresInside(grid [][]int) int {
	cnt := 0
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[i])-2; j++ {
			if checkMagic(grid, i, j) {
				cnt++
			}
		}
	}
	return cnt
}

// i, j is top, left pos
func checkMagic(grid [][]int, i, j int) bool {
	if !from1to9(grid, i, j) {
		return false
	}
	// 1 cross
	a := grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2]
	for k := i; k <= i+2; k++ {
		if a != grid[k][j]+grid[k][j+1]+grid[k][j+2] {
			return false
		}
	}
	for k := j; k <= j+2; k++ {
		if a != grid[i][k]+grid[i+1][k]+grid[i+2][k] {
			return false
		}
	}
	// This case will never happen
	// if a != grid[i+2][j]+grid[i+1][j+1]+grid[i][j+2] {
	// 	return false
	// }
	return true
}

func from1to9(grid [][]int, i, j int) bool {
	cnt := make([]int, 10)
	for x := range 3 {
		for y := range 3 {
			if grid[i+x][j+y] == 0 || grid[i+x][j+y] > 9 {
				return false
			}
			if cnt[grid[i+x][j+y]] > 0 {
				return false
			}
			cnt[grid[i+x][j+y]]++
		}
	}
	return true
}
