package weekly397

import (
	"math"

	"github.com/lovung/ds/matrix"
	"github.com/lovung/ds/queue"
)

// https://leetcode.com/problems/maximum-difference-score-in-a-grid/description/
func maxScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := matrix.New[int](m, n)
	res := math.MinInt
	for i := range m {
		for j := range n {
			val := math.MaxInt
			if matrix.In(grid, i-1, j) {
				val = min(val, dp.At(i-1, j))
			}
			if matrix.In(grid, i, j-1) {
				val = min(val, dp.At(i, j-1))
			}
			res = max(res, grid[i][j]-val)
			dp.UpdateAt(i, j, min(val, grid[i][j]))
		}
	}
	return res
}

func maxScore2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m+1)
	res := math.MinInt
	for i := range m + 1 {
		dp[i] = make([]int, n+1)
		dp[i][0] = math.MaxInt
	}
	for j := range dp[0] {
		dp[0][j] = math.MaxInt
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			val := min(dp[i][j+1], dp[i+1][j])
			if i > 0 {
				val = min(val, grid[i-1][j])
			}
			if j > 0 {
				val = min(val, grid[i][j-1])
			}
			dp[i+1][j+1] = val
			res = max(res, grid[i][j]-dp[i+1][j+1])
		}
	}
	return res
}

// Wrong solution because of we only go with positive score during BFS
func maxScore1(grid [][]int) int {
	type pos struct{ x, y int }
	type cell struct{ x, y, score int }
	m, n := len(grid), len(grid[0])
	q := queue.NewCircularQueue[cell](1e5)
	res := math.MinInt
	for i := m - 1; i >= 0; i-- {
		maxValInRow := grid[i][n-1]
		for j := n - 2; j >= 0; j-- {
			scoreIfJump := maxValInRow - grid[i][j]
			if scoreIfJump > 0 {
				q.EnQueue(cell{i, j, scoreIfJump})
			}
			maxValInRow = max(maxValInRow, grid[i][j])
			res = max(res, scoreIfJump)
		}
	}
	for j := n - 1; j >= 0; j-- {
		maxValInCol := grid[m-1][j]
		for i := m - 2; i >= 0; i-- {
			scoreIfJump := maxValInCol - grid[i][j]
			if scoreIfJump > 0 {
				q.EnQueue(cell{i, j, scoreIfJump})
			}
			maxValInCol = max(maxValInCol, grid[i][j])
			res = max(res, scoreIfJump)
		}
	}
	cache := make(map[pos]int)
	for q.Len() > 0 {
		qCell, _ := q.DeQueue()
		res = max(res, qCell.score)
		cache[pos{qCell.x, qCell.y}] = qCell.score
		for i := qCell.x - 1; i >= 0; i-- {
			if score := grid[qCell.x][qCell.y] - grid[i][qCell.y]; score > 0 {
				if qCell.score+score > cache[pos{qCell.x, qCell.y}] {
					q.EnQueue(cell{i, qCell.y, qCell.score + score})
				}
			}
		}
		for j := qCell.y - 1; j >= 0; j-- {
			if score := grid[qCell.x][qCell.y] - grid[qCell.x][j]; score > 0 {
				if qCell.score+score > cache[pos{qCell.x, qCell.y}] {
					q.EnQueue(cell{qCell.x, j, qCell.score + score})
				}
			}
		}
	}
	return res
}
