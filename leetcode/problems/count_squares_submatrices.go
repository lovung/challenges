package problems

// Link: https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3336/
func countSquares(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	out := 0

	for i := 0; i < n; i++ {
		out += matrix[m-1][i]
	}
	for i := 0; i < m; i++ {
		out += matrix[i][n-1]
	}
	out -= matrix[m-1][n-1]
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			if matrix[i][j] == 1 {
				matrix[i][j] = 1 + min(matrix[i+1][j+1], min(matrix[i][j+1], matrix[i+1][j]))
			}
			out += matrix[i][j]
		}
	}
	return out
}
