package contest

import (
	"math"

	"github.com/lovung/ds/matrix"
	"github.com/lovung/ds/queue"
)

type pos struct {
	i, j int
}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	mat := matrix.NewSquare[int](n)
	qThiefPos := queue.NewSimpleQueue[pos]()
	mat.ForEach(func(_, i, j int) {
		if grid[i][j] == 0 {
			mat.UpdateAt(i, j, math.MaxInt64)
		} else {
			mat.UpdateAt(i, j, 1)
			qThiefPos.EnQueue(pos{i, j})
		}
	})

	ans := math.MaxInt64
	level := 2
	for qThiefPos.Len() > 0 {
		for n := qThiefPos.Len(); n > 0; n-- {
			floodPos, _ := qThiefPos.DeQueue()
			mat.ForEachNearBy(floodPos.i, floodPos.j, func(val, x, y int) {
				if val > level {
					mat.UpdateAt(x, y, level)
					qThiefPos.EnQueue(pos{x, y})
				}
			})
		}
		level++
	}

	l, h := 0, level-1
	for l <= h {
		m := (l + h) / 2
		if isPathAbleToExist(mat, m) {
			ans, l = m, m+1
		} else {
			h = m - 1
		}
	}
	return ans - 1
}

func isPathAbleToExist(mat *matrix.Matrix[int], lowestLevel int) bool {
	n := len(mat.Data())
	vis := matrix.NewSquare[bool](n)
	bfsQueue := queue.NewSimpleQueue[pos]()
	if mat.At(0, 0) >= lowestLevel {
		bfsQueue.EnQueue(pos{0, 0})
	}
	vis.UpdateAt(0, 0, true)
	for bfsQueue.Len() > 0 {
		bfsPos, _ := bfsQueue.DeQueue()
		if bfsPos.i == n-1 && bfsPos.j == n-1 {
			return true
		}
		mat.ForEachNearBy(bfsPos.i, bfsPos.j, func(val, x, y int) {
			if !vis.At(x, y) && val >= lowestLevel {
				vis.UpdateAt(x, y, true)
				bfsQueue.EnQueue(pos{x, y})
			}
		})
	}
	return false
}
