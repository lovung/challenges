package may2024

// https://leetcode.com/problems/largest-local-values-in-a-matrix
func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	res := make([][]int, n-2)
	for i := range res {
		res[i] = make([]int, n-2)
	}
	for i := range n - 2 {
		for j := range n - 2 {
			res[i][j] = maxLocal(grid, i, j)
		}
	}
	return res
}

var cells = [][]int{
	{0, 0}, {0, 1}, {0, 2},
	{1, 0}, {1, 1}, {1, 2},
	{2, 0}, {2, 1}, {2, 2},
}

func maxLocal(grid [][]int, i, j int) int {
	res := 0
	for _, c := range cells {
		res = max(res, grid[i+c[0]][j+c[1]])
	}
	return res
}
