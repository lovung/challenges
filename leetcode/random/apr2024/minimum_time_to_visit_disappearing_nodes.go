package apr2024

import (
	"container/heap"

	"github.com/lovung/ds/heaps"
)

type dijkstraNode struct {
	src   int
	point int
}

func (i dijkstraNode) Int() int {
	return i.point
}

// https://leetcode.com/problems/minimum-time-to-visit-disappearing-nodes/
func minimumTime(n int, edges [][]int, disappear []int) []int {
	adjMap := make(map[int]map[int]int)
	for _, e := range edges {
		if adjMap[e[0]] == nil {
			adjMap[e[0]] = make(map[int]int)
		}
		if adjMap[e[0]][e[1]] == 0 {
			adjMap[e[0]][e[1]] = e[2]
		} else {
			adjMap[e[0]][e[1]] = min(adjMap[e[0]][e[1]], e[2])
		}
		if adjMap[e[1]] == nil {
			adjMap[e[1]] = make(map[int]int)
		}
		if adjMap[e[1]][e[0]] == 0 {
			adjMap[e[1]][e[0]] = e[2]
		} else {
			adjMap[e[1]][e[0]] = min(adjMap[e[1]][e[0]], e[2])
		}
	}

	dijk := make([]int, n)
	for i := range dijk {
		dijk[i] = -1
	}
	pq := heaps.MinHeapInt[dijkstraNode]{}
	heap.Init(&pq)
	heap.Push(&pq, dijkstraNode{0, 0})
	for pq.Len() > 0 {
		curNode := heap.Pop(&pq).(dijkstraNode)
		if dijk[curNode.src] != -1 {
			continue
		}
		dijk[curNode.src] = curNode.point
		for k, v := range adjMap[curNode.src] {
			newV := curNode.point + v
			if (dijk[k] != -1 && newV >= dijk[k]) || newV >= disappear[k] {
				continue
			}
			heap.Push(&pq, dijkstraNode{k, newV})
		}
	}
	return dijk
}
