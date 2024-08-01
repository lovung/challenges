package jul2024

// https://leetcode.com/problems/lucky-numbers-in-a-matrix/
func luckyNumbers(matrix [][]int) []int {
	minInRows := make([]int, len(matrix))
	maxInCols := make([]int, len(matrix[0]))
	for i := range matrix {
		minInRow := int(1e5)
		for j := range matrix[i] {
			minInRow = min(minInRow, matrix[i][j])
		}
		minInRows[i] = minInRow
	}
	for j := range matrix[0] {
		maxInCol := 0
		for i := range matrix {
			maxInCol = max(maxInCol, matrix[i][j])
		}
		maxInCols[j] = maxInCol
	}
	var res []int
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == minInRows[i] &&
				matrix[i][j] == maxInCols[j] {
				res = append(res, matrix[i][j])
			}
		}
	}
	return res
}
