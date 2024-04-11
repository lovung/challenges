package apr2024

// More cases
// 4. [["A","B","C","E"],["S","F","E","S"],["A","D","E","E"]]
// "ABCESEEEFS"
// 5. [["A","A","A","A","A","A"],["A","A","A","A","A","A"],["A","A","A","A","A","A"],["A","A","A","A","A","A"],["A","A","A","A","A","A"],["A","A","A","A","A","A"]]
// "AAAAAAAAAAAAAAB"

var dirs = [][]int{
	{1, 0}, {0, 1}, {0, -1}, {-1, 0},
}

type xy struct {
	x, y int
}

func countChar(board [][]byte, word string) bool {
	count := make(map[byte]int)
	for _, v := range board {
		for _, j := range v {
			count[j] += 1
		}
	}

	for _, char := range word {
		find, _ := count[byte(char)]
		if find < 1 {
			return false
		}
		count[byte(char)] -= 1
	}
	return true
}

func exist(board [][]byte, word string) bool {
	if !countChar(board, word) {
		return false
	}

	wordBytes := []byte(word)
	for i := range board {
		for j := range board[i] {
			if board[i][j] != word[0] {
				continue
			}
			visited := [6][6]bool{}
			if dfs(board, wordBytes[1:], i, j, &visited) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, remain []byte, x, y int, visited *[6][6]bool) bool {
	if len(remain) == 0 {
		return true
	}
	(*visited)[x][y] = true
	for _, dir := range dirs {
		if !inBound(board, x+dir[0], y+dir[1]) {
			continue
		}
		if (*visited)[x+dir[0]][y+dir[1]] {
			continue
		}
		if board[x+dir[0]][y+dir[1]] != remain[0] {
			continue
		}
		if dfs(board, remain[1:], x+dir[0], y+dir[1], visited) {
			return true
		}
	}
	(*visited)[x][y] = false
	return false
}

func inBound[T any](mat [][]T, i, j int) bool {
	if i < 0 || j < 0 || i >= len(mat) || j >= len(mat[0]) {
		return false
	}
	return true
}
