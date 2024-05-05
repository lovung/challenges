package apr2024

// https://leetcode.com/problems/maximal-rectangle/
func maximalRectangle3(matrix [][]byte) int {
	n, m := len(matrix), len(matrix[0])
	suffix := make([][]int, n+1)
	for i := range suffix {
		suffix[i] = make([]int, len(matrix[0])+1)
	}
	for j := m - 1; j >= 0; j-- {
		for i := n - 1; i >= 0; i-- {
			if matrix[i][j] == '1' {
				suffix[i][j] = suffix[i+1][j] + 1
			}
		}
	}
	res := 0
	for i := 0; i < len(matrix); i++ {
		l := 0
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				l = j + 1
				continue
			}
			curMin := suffix[i][j]
			for k := j; k >= l; k-- {
				curMin = min(curMin, suffix[i][k])
				area := curMin * (j - k + 1)
				res = max(res, area)
			}
		}
	}
	return res
}
