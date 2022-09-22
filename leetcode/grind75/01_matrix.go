package grind75

// Link: https://leetcode.com/problems/01-matrix/
var dir = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func updateMatrix(mat [][]int) [][]int {
	queue := make([][]int, 0)
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}
	for len(queue) > 0 {
		first := queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = make([][]int, 0)
		}
		for _, d := range dir {
			i, j := first[0]+d[0], first[1]+d[1]
			if !inBound(mat, i, j) {
				continue
			}
			if mat[i][j] != -1 {
				continue
			}
			mat[i][j] = mat[first[0]][first[1]] + 1
			queue = append(queue, []int{i, j})
		}
	}
	return mat
}

func inBound(mat [][]int, i, j int) bool {
	if i < 0 || j < 0 || i >= len(mat) || j >= len(mat[0]) {
		return false
	}
	return true
}
