package may2024

import (
	"container/heap"
	"math"

	"github.com/lovung/ds/heaps"
	"github.com/lovung/ds/matrix"
	"github.com/lovung/ds/queue"
)

type pos struct{ x, y int }

// https://leetcode.com/problems/find-the-safest-path-in-a-grid/
func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	thiefs := make([]pos, 0, n)
	safeness := make([][]int, n)
	for i := range safeness {
		safeness[i] = make([]int, n)
		for j := range safeness {
			if grid[i][j] == 0 {
				safeness[i][j] = math.MaxInt
			} else {
				thiefs = append(thiefs, pos{i, j})
			}
		}
	}
	for _, t := range thiefs {
		floodfill(safeness, t)
	}
	return dijkstra(grid, safeness)
}

type dijkstraNode struct {
	pos      pos
	safeness int
}

func dijkstra(grid [][]int, safeness [][]int) int {
	if safeness[0][0] == 0 {
		return 0
	}
	n := len(grid)
	pq := heaps.MaxHeapWithValue[int]{}
	heap.Init(&pq)
	item := heaps.HeapItem[int]{
		Ref:   safeness[0][0],
		Value: pos{0, 0},
	}
	heap.Push(&pq, &item)
	res := safeness[0][0]
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*heaps.HeapItem[int])
		ix := item.Value.(pos).x
		iy := item.Value.(pos).y
		visited[ix][iy] = true
		res = min(res, item.Ref)
		if ix == n-1 && iy == n-1 {
			break
		}
		for _, d := range matrix.Dirs {
			x := ix + d[0]
			y := iy + d[1]
			if !matrix.In(safeness, x, y) || visited[x][y] {
				continue
			}
			heap.Push(&pq, &heaps.HeapItem[int]{
				Ref:   safeness[x][y],
				Value: pos{x, y},
			})
		}
	}
	return res
}

func floodfill(safeness [][]int, start pos) {
	n := len(safeness)
	q := queue.NewCircularQueue[pos](n * n)
	q.EnQueue(start)
	for q.Len() > 0 {
		item, _ := q.DeQueue()
		for _, d := range matrix.Dirs {
			x := item.x + d[0]
			y := item.y + d[1]
			if !matrix.In(safeness, x, y) {
				continue
			}
			if safeness[x][y] <= safeness[item.x][item.y]+1 {
				continue
			}
			safeness[x][y] = safeness[item.x][item.y] + 1
			q.EnQueue(pos{x, y})
		}
	}
}

func maximumSafenessFactor2(grid [][]int) int {
	n := len(grid)
	mat := matrix.NewSquare[int](n)
	qThiefPos := queue.NewSimpleQueue[pos]()
	mat.ForEach(func(_, i, j int) {
		if grid[i][j] == 0 {
			mat.UpdateAt(i, j, math.MaxInt64)
		} else {
			mat.UpdateAt(i, j, 0)
			qThiefPos.EnQueue(pos{i, j})
		}
	})

	ans := math.MaxInt64
	level := 1
	for qThiefPos.Len() > 0 {
		for n := qThiefPos.Len(); n > 0; n-- {
			floodPos, _ := qThiefPos.DeQueue()
			mat.ForEachNearBy(floodPos.x, floodPos.y, func(val, x, y int) {
				if val > level {
					mat.UpdateAt(x, y, level)
					qThiefPos.EnQueue(pos{x, y})
				}
			})
		}
		level++
	}

	l, h := 0, level
	for l <= h {
		m := (l + h) / 2
		if isPathAbleToExist(mat, m) {
			ans, l = m, m+1
		} else {
			h = m - 1
		}
	}
	return ans
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
		if bfsPos.x == n-1 && bfsPos.y == n-1 {
			return true
		}
		mat.ForEachNearBy(bfsPos.x, bfsPos.y, func(val, x, y int) {
			if !vis.At(x, y) && val >= lowestLevel {
				vis.UpdateAt(x, y, true)
				bfsQueue.EnQueue(pos{x, y})
			}
		})
	}
	return false
}
