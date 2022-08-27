package medium

import "github.com/lovung/ds/graphs"

// Link: https://leetcode.com/problems/surrounded-regions/
const x, o = 'X', 'O'

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board[0]), len(board)
	type point struct{ x, y int }
	mark := make([][]bool, n)
	for i := 0; i < n; i++ {
		mark[i] = make([]bool, m)
	}
	// q := queue.NewQueue[*point]()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
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

// Ref: https://github.com/halfrost/LeetCode-Go/blob/master/leetcode/0130.Surrounded-Regions/130.%20Surrounded%20Regions.go
func solveWithUnionFind(board [][]byte) {
	if len(board) == 0 {
		return
	}
	m, n := len(board[0]), len(board)
	uf := graphs.NewUnionFind(n*m + 1)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (i == 0 || i == n-1 || j == 0 || j == m-1) && board[i][j] == o { //棋盘边缘上的 o 点
				uf.Union(i*m+j, n*m)
			} else if board[i][j] == o {
				if board[i-1][j] == o {
					uf.Union(i*m+j, (i-1)*m+j)
				}
				if board[i+1][j] == o {
					uf.Union(i*m+j, (i+1)*m+j)
				}
				if board[i][j-1] == o {
					uf.Union(i*m+j, i*m+j-1)
				}
				if board[i][j+1] == o {
					uf.Union(i*m+j, i*m+j+1)
				}

			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if uf.Find(i*m+j) != uf.Find(n*m) {
				board[i][j] = x
			}
		}
	}
}
