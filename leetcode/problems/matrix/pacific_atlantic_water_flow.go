package matrix

// Link: https://leetcode.com/problems/pacific-atlantic-water-flow/
func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	pacificMark := make([][]bool, m)
	atlanticMark := make([][]bool, m)
	for i := 0; i < m; i++ {
		pacificMark[i] = make([]bool, n)
		atlanticMark[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		if i == 0 {
			for j := 0; j < n; j++ {
				floodField(heights, pacificMark, 0, i, j)
			}
		} else {
			floodField(heights, pacificMark, 0, i, 0)
		}
	}
	for i := 0; i < m; i++ {
		if i == m-1 {
			for j := 0; j < n; j++ {
				floodField(heights, atlanticMark, 0, i, j)
			}
		} else {
			floodField(heights, atlanticMark, 0, i, n-1)
		}
	}
	res := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacificMark[i][j] && atlanticMark[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func floodField(heights [][]int, mark [][]bool, lastHeight, i, j int) {
	if i < 0 || j < 0 || i >= len(heights) || j >= len(heights[0]) {
		return
	}
	if mark[i][j] {
		return
	}
	if lastHeight > heights[i][j] {
		return
	}
	mark[i][j] = true
	floodField(heights, mark, heights[i][j], i-1, j)
	floodField(heights, mark, heights[i][j], i, j-1)
	floodField(heights, mark, heights[i][j], i+1, j)
	floodField(heights, mark, heights[i][j], i, j+1)
}
