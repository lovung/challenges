package weekly403

// https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-i/description/
func minimumArea(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	return getMinArea(grid, 0, n-1, 0, m-1)
}

func getMinArea(grid [][]int, fx, tx, fy, ty int) int {
	n, m := len(grid), len(grid[0])
	tlX, tlY, brX, brY := n, m, 0, 0
	for x := fx; x <= tx; x++ {
		for y := fy; y <= ty; y++ {
			if grid[x][y] == 1 {
				tlX = min(tlX, x)
				tlY = min(tlY, y)
				brX = max(brX, x)
				brY = max(brY, y)
			}
		}
	}
	return (brX - tlX + 1) * (brY - tlY + 1)
}
