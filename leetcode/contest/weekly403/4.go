package weekly403

// https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-ii/description/
func minimumSum(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	ans := m * n
	for i := range n - 2 {
		for j := i + 1; j < n-1; j++ {
			ans = min(ans,
				getMinArea(grid, 0, i, 0, m-1)+
					getMinArea(grid, i+1, j, 0, m-1)+
					getMinArea(grid, j+1, n-1, 0, m-1))
		}
	}
	for i := range m - 2 {
		for j := i + 1; j < m-1; j++ {
			ans = min(ans,
				getMinArea(grid, 0, n-1, 0, i)+
					getMinArea(grid, 0, n-1, i+1, j)+
					getMinArea(grid, 0, n-1, j+1, m-1))
		}
	}
	for i := range n - 1 {
		for j := range m - 1 {
			ans = min(ans,
				getMinArea(grid, 0, n-1, 0, j)+
					getMinArea(grid, 0, i, j+1, m-1)+
					getMinArea(grid, i+1, n-1, j+1, m-1))

			ans = min(ans,
				getMinArea(grid, 0, i, 0, j)+
					getMinArea(grid, i+1, n-1, 0, j)+
					getMinArea(grid, 0, n-1, j+1, m-1))

			ans = min(ans,
				getMinArea(grid, 0, i, 0, j)+
					getMinArea(grid, 0, i, j+1, m-1)+
					getMinArea(grid, i+1, n-1, 0, m-1))

			ans = min(ans,
				getMinArea(grid, 0, i, 0, m-1)+
					getMinArea(grid, i+1, n-1, 0, j)+
					getMinArea(grid, i+1, n-1, j+1, m-1))
		}
	}
	return ans
}
