package biweekly129

// https://leetcode.com/contest/biweekly-contest-129/problems/make-a-square-with-the-same-color/
func canMakeSquare(grid [][]byte) bool {
	const (
		b = 'B'
	)
	type pair struct {
		x, y int
	}
	c := [][]pair{
		{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		{{0, 1}, {1, 1}, {0, 2}, {1, 2}},
		{{1, 0}, {1, 1}, {2, 0}, {2, 1}},
		{{1, 1}, {2, 1}, {1, 2}, {2, 2}},
	}
	for i := range c {
		cntB := 0
		for _, p := range c[i] {
			if grid[p.x][p.y] == b {
				cntB++
			}
		}
		if cntB != 2 {
			return true
		}
	}
	return false
}
