package oct2022

// Link: https://leetcode.com/problems/valid-tic-tac-toe-state/
func validTicTacToe(board []string) bool {
	cntX, cntO := 0, 0
	// count
	for i := range board {
		for j := range board[i] {
			switch board[i][j] {
			case 'X':
				cntX++
			case 'O':
				cntO++
			}
		}
	}
	if cntO != cntX && cntO != cntX-1 {
		return false
	}
	// check if game is over
	who, isWin := checkWin(board)
	if !isWin {
		return true
	}
	// if O win, must cntO == cntX
	if who == 'O' {
		return cntO == cntX
	}

	// if X win, must cntO == cntX-1
	return cntO == cntX-1
}

var groupCheck = [][]int{
	{0, 0}, {0, 1}, {0, 2},
	{1, 0}, {1, 1}, {1, 2},
	{2, 0}, {2, 1}, {2, 2},
	{0, 0}, {1, 0}, {2, 0},
	{0, 1}, {1, 1}, {2, 1},
	{0, 2}, {1, 2}, {2, 2},
	{0, 0}, {1, 1}, {2, 2},
	{0, 2}, {1, 1}, {2, 0},
}

func checkWin(board []string) (who byte, win bool) {
	for i := 0; i < len(groupCheck); i += 3 {
		ref := board[groupCheck[i][0]][groupCheck[i][1]]
		if ref == ' ' {
			continue
		}
		if ref == board[groupCheck[i+1][0]][groupCheck[i+1][1]] &&
			board[groupCheck[i+2][0]][groupCheck[i+2][1]] == ref {
			return ref, true
		}
	}
	return 0, false
}
