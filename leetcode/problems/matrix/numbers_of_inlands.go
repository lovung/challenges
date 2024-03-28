package matrix

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

// Link: https://leetcode.com/problems/number-of-islands/submissions/
func numIslands(grid [][]byte) int {
	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				searchIslands(grid, i, j)
				res++
			}
		}
	}
	return res
}

func searchIslands(grid [][]byte, x, y int) {
	grid[x][y] = '2'
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		if isInBoard(grid, nx, ny) && grid[nx][ny] == '1' {
			searchIslands(grid, nx, ny)
		}
	}
}

func isInBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}
