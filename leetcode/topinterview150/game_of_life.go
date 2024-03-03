package topinterview150

var dirs = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, -1},
	{-1, 1},
}

const (
	live = 1
	dead = 0
)

// Link: https://leetcode.com/problems/game-of-life/description/
func gameOfLife(board [][]int) {
	clone := make([][]int, 0, len(board))
	for i := range board {
		newRow := make([]int, len(board[i]))
		copy(newRow, board[i])
		clone = append(clone, newRow)
	}
	for i := range clone {
		for j := range clone[i] {
			liveCnt := countLiveNeighbors(clone, i, j)
			switch {
			case clone[i][j] == dead && liveCnt == 3:
				board[i][j] = live
			case clone[i][j] == live && (liveCnt < 2 || liveCnt > 3):
				board[i][j] = dead
			}
		}
	}
}

func within(board [][]int, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && (len(board) == 0 || y < len(board[0]))
}

func countLiveNeighbors(board [][]int, x, y int) int {
	count := 0
	for _, dir := range dirs {
		nx, ny := x+dir[0], y+dir[1]
		if within(board, nx, ny) && board[nx][ny] == live {
			count++
		}
	}
	return count
}
