package may2024

// https://leetcode.com/problems/score-after-flipping-matrix/
func matrixScore(grid [][]int) int {
	for i := range grid {
		if grid[i][0] == 0 {
			flipRow(grid, i)
		}
	}
	m := len(grid)
	for j := range len(grid[0]) {
		if countZero(grid, j) > m/2 {
			flipCol(grid, j)
		}
	}
	return sum(grid)
}

func sum(grid [][]int) int {
	res := 0
	for i := range grid {
		num := 0
		bit := 1
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				num |= bit
			}
			bit <<= 1
		}
		res += num
	}
	return res
}

func flipRow(grid [][]int, r int) {
	for c := range grid[r] {
		grid[r][c] ^= 1
	}
}

func flipCol(grid [][]int, c int) {
	for r := range grid {
		grid[r][c] ^= 1
	}
}
func countZero(grid [][]int, c int) int {
	cnt := 0
	for r := range grid {
		if grid[r][c] == 0 {
			cnt++
		}
	}
	return cnt
}
