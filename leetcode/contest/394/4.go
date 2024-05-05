package contest394

import (
	"container/heap"

	"github.com/lovung/ds/heaps"
)

type dijkstraWalk struct {
	src  int
	cost int
	path []int
}

func (i dijkstraWalk) Int() int {
	return i.cost
}

type dijkNode struct {
	cost  *int
	paths [][]int
}

type edge struct {
	i, w int
}

func findAnswer(n int, edges [][]int) []bool {
	adjMap := make(map[int]map[int]edge)
	for i, e := range edges {
		if adjMap[e[0]] == nil {
			adjMap[e[0]] = make(map[int]edge)
		}
		adjMap[e[0]][e[1]] = edge{i, e[2]}
		if adjMap[e[1]] == nil {
			adjMap[e[1]] = make(map[int]edge)
		}
		adjMap[e[1]][e[0]] = edge{i, e[2]}
	}

	dijk := make([]dijkNode, n)
	pq := heaps.MinHeapInt[dijkstraWalk]{}
	heap.Init(&pq)
	heap.Push(&pq, dijkstraWalk{0, 0, nil})
	for pq.Len() > 0 {
		walk := heap.Pop(&pq).(dijkstraWalk)
		if dijk[walk.src].cost != nil {
			if walk.cost > *dijk[walk.src].cost {
				continue
			}
			dijk[walk.src].paths = append(dijk[walk.src].paths, walk.path)
		} else {
			dijk[walk.src].paths = [][]int{walk.path}
		}
		dijk[walk.src].cost = &walk.cost
		for k, v := range adjMap[walk.src] {
			newV := walk.cost + v.w
			newPath := make([]int, len(walk.path)+1)
			copy(newPath, walk.path)
			newPath[len(newPath)-1] = v.i
			heap.Push(&pq, dijkstraWalk{k, newV, newPath})
		}
	}
	res := make([]bool, len(edges))
	for _, p := range dijk[n-1].paths {
		for _, n := range p {
			res[n] = true
		}
	}
	return res
}
