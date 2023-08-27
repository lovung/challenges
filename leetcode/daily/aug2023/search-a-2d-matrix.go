package aug2023

// https://leetcode.com/problems/search-a-2d-matrix/
func searchMatrix(matrix [][]int, target int) bool {
	rowIdx := binarySearchRow(matrix, target)
	if rowIdx < 0 {
		return false
	}
	col := binarySearchColInRow(matrix[rowIdx], target)
	return col >= 0 && col < len(matrix[0])
}

func binarySearchRow(matrix [][]int, target int) int {
	if target < matrix[0][0] {
		return -1
	}
	l, r := 0, len(matrix)-1
	for l < r {
		if target == matrix[l][0] {
			return l
		}
		if target >= matrix[r][0] {
			return r
		}
		m := (l + r) / 2
		if l == m {
			return l
		}
		if target < matrix[m][0] {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}

func binarySearchColInRow(row []int, target int) int {
	l, r := 0, len(row)-1
	for l < r {
		m := (l + r) / 2
		if target == row[m] {
			return m
		}
		if target < row[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	if row[l] == target {
		return l
	}
	return -1
}
