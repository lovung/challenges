package grind75

// Link: https://leetcode.com/problems/spiral-matrix/
func spiralOrder(matrix [][]int) []int {
	cnt := len(matrix) * len(matrix[0])
	curDir := 0

	x, y := 0, 0
	res := make([]int, 0, cnt)
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}
	for i := 0; i < cnt; i++ {
		res = append(res, matrix[x][y])
		visited[x][y] = true
		x, y = calNextPoint(matrix, x, y, &curDir, visited)
	}
	return res
}

func calNextPoint(matrix [][]int, x, y int, curDir *int, visited [][]bool) (int, int) {
	if !inBound(matrix, x+dir[*curDir][0], y+dir[*curDir][1]) || visited[x+dir[*curDir][0]][y+dir[*curDir][1]] {
		*curDir = (*curDir + 1) % 4 // rotate
	}
	return x + dir[*curDir][0], y + dir[*curDir][1]
}
