package contest392

import (
	"math"

	"github.com/lovung/ds/graphs"
)

type UnionFind struct {
	graphs.UnionFind
	weight []int
}

// NewUnionFind init the UnionFind with size
func NewUnionFind(size int) UnionFind {
	uf := UnionFind{
		UnionFind: graphs.NewUnionFind(size),
		weight:    make([]int, size),
	}
	for i := 0; i < size; i++ {
		uf.weight[i] = math.MaxInt64
	}
	return uf
}

func (uf *UnionFind) AddWeight(x, w int) {
	p := uf.Find(x)
	uf.weight[p] &= w
}

func (uf *UnionFind) Weight(x int) int {
	p := uf.Find(x)
	if uf.weight[p] == math.MaxInt64 {
		return -1
	}
	return uf.weight[p]
}
func minimumCost(n int, edges [][]int, query [][]int) []int {
	cuf := NewUnionFind(n)
	for _, e := range edges {
		cuf.Union(e[0], e[1])
	}
	for _, e := range edges {
		cuf.AddWeight(e[0], e[2])
	}
	res := make([]int, 0, len(query))
	for _, q := range query {
		val := 0
		if q[0] != q[1] {
			p := cuf.Find(q[0])
			val = cuf.Weight(p)
			if p != cuf.Find(q[1]) {
				val = -1
			}
		}
		res = append(res, val)
	}
	return res
}
