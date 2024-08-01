package jul2024

// https://leetcode.com/problems/find-valid-matrix-given-row-and-column-sums/
func restoreMatrix(rowSum []int, colSum []int) [][]int {
	matrix := make([][]int, len(rowSum))
	for i := range matrix {
		matrix[i] = make([]int, len(colSum))
	}
	for i := range matrix {
		for j := range matrix[i] {
			x := min(rowSum[i], colSum[j])
			matrix[i][j] = x
			rowSum[i] -= x
			colSum[j] -= x
		}
	}
	return matrix
}
