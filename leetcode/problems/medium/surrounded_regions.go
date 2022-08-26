package medium

// Link: https://leetcode.com/problems/surrounded-regions/
const x, o = 'X', 'O'

func solve(board [][]byte) {
	type point struct{ x, y int }
	mark := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		mark[i] = make([]bool, len(board[i]))
	}
	// q := queue.NewQueue[*point]()
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == x {
				continue
			}
			clearMark(mark)
			needTurn := make([][]int, 0)
			if solveBFSIsSurround(board, mark, i, j, &needTurn) {
				turnToX(board, needTurn)
			}
		}
	}
}

func turnToX(board [][]byte, needTurn [][]int) {
	for i := range needTurn {
		board[needTurn[i][0]][needTurn[i][1]] = x
	}
}

func clearMark(mark [][]bool) {
	for i := 0; i < len(mark); i++ {
		for j := 0; j < len(mark[0]); j++ {
			mark[i][j] = false
		}
	}
}

func solveBFSIsSurround(board [][]byte, mark [][]bool, i, j int, needTurn *[][]int) bool {
	if mark[i][j] {
		return true
	}
	mark[i][j] = true
	if board[i][j] == x {
		return true
	}
	if i == 0 || j == 0 || i == len(board)-1 || j == len(board[0])-1 {
		return false
	}
	*needTurn = append(*needTurn, []int{i, j})
	return solveBFSIsSurround(board, mark, i-1, j, needTurn) &&
		solveBFSIsSurround(board, mark, i+1, j, needTurn) &&
		solveBFSIsSurround(board, mark, i, j-1, needTurn) &&
		solveBFSIsSurround(board, mark, i, j+1, needTurn)
}
