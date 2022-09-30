package hard

import (
	"container/heap"
	"sort"

	"github.com/lovung/ds/heaps"
)

type point struct{ i, h int }

// Link: https://leetcode.com/problems/the-skyline-problem/
func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	ps := make([]point, 0, n*2)
	for _, v := range buildings {
		ps = append(ps, point{v[0], -v[2]}, point{v[1], v[2]})
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].i < ps[j].i || ps[i].i == ps[j].i && ps[i].h < ps[j].h
	})

	maxHeap := &heaps.MaxHeap[int]{}
	heap.Push(maxHeap, 0)
	prev := 0
	res := [][]int{}
	for _, p := range ps {
		if p.h < 0 {
			heap.Push(maxHeap, -p.h)
		} else {
			heap.Remove(maxHeap, maxHeap.Index(p.h))
		}
		curr := (*maxHeap)[0]
		if curr != prev {
			res = append(res, []int{p.i, curr})
			prev = curr
		}
	}
	return res
}
