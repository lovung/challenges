package grind75

import (
	"github.com/lovung/ds/maths"
	"github.com/lovung/ds/queue"
)

// Link: https://leetcode.com/problems/rotting-oranges/
type pointWithDis struct{ i, j, dis int }

type point struct{ i, j int }

func orangesRotting(grid [][]int) int {
	cacheDistance := make(map[point]int)

	q := queue.NewSimpleQueue[pointWithDis]()
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				q.EnQueue(pointWithDis{i, j, 0})
			} else if grid[i][j] == 1 {
				cacheDistance[point{i, j}] = -1
			}
		}
	}

	if q.Len() == 0 && len(cacheDistance) > 0 {
		return -1
	}

	for q.Len() > 0 {
		orange, _ := q.DeQueue()
		floody(grid, cacheDistance, q, orange.i, orange.j, orange.dis)
	}
	maxDis := 0
	for _, v := range cacheDistance {
		if v == -1 {
			return -1
		}
		maxDis = maths.Max(maxDis, v)
	}
	return maxDis
}

func floody(
	grid [][]int, cache map[point]int,
	q queue.Queue[pointWithDis],
	i, j int, dis int,
) {
	for _, d := range dir {
		_i := i + d[0]
		_j := j + d[1]
		if !inBound(grid, _i, _j) {
			continue
		}
		if grid[_i][_j] == 2 {
			continue
		}
		if grid[_i][_j] == 1 {
			if cache[point{_i, _j}] == -1 {
				cache[point{_i, _j}] = dis + 1
				q.EnQueue(pointWithDis{_i, _j, cache[point{_i, _j}]})
			} else {
				cache[point{_i, _j}] = maths.Min(cache[point{_i, _j}], dis+1)
			}
		}
	}
}
