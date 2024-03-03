package microsoft

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

const (
	land        = 1
	visitedLand = 2
)

// Link: https://leetcode.com/problems/max-area-of-island/description/
func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == land {
				area := 0
				dfsScanIslands(grid, i, j, &area)
				res = max(res, area)
			}
		}
	}
	return res
}

func dfsScanIslands(grid [][]int, x, y int, area *int) {
	grid[x][y] = visitedLand
	*area++
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		if isInBoard(grid, nx, ny) && grid[nx][ny] == land {
			dfsScanIslands(grid, nx, ny, area)
		}
	}
}

func isInBoard(board [][]int, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}
