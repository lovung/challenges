package jun2024

import (
	"container/heap"
	"slices"

	"github.com/lovung/ds/trees"
)

type project struct{ capital, profit, index int }

func projectProfitCmp(n1, n2 *project) *project {
	if n1 == nil && n2 != nil {
		return n2
	}
	if n1 != nil && n2 == nil {
		return n1
	}
	if n1.profit > n2.profit {
		return n1
	}
	return n2
}

// https://leetcode.com/problems/ipo
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	nodes := make([]*project, 0, n)
	for i := range n {
		nodes = append(nodes, &project{capital[i], profits[i], i})
	}
	slices.SortFunc(nodes, func(a, b *project) int {
		return a.capital - b.capital
	})
	st := trees.NewSegmentTreeWithArray(n, projectProfitCmp)
	for i := range n {
		nodes[i].index = i
		st.Update(i, nodes[i])
	}
	for range k {
		i, ok := slices.BinarySearchFunc(nodes, w, func(n *project, c int) int { return n.capital - c })
		for ok && i < n && nodes[i].capital == w {
			i++
		}
		if i == 0 {
			break
		}
		maxProfitNode := st.Query(0, i)
		w += maxProfitNode.profit
		maxProfitNode.profit = 0
		st.Update(maxProfitNode.index, maxProfitNode)
	}
	return w
}

type PausedHeap []project
type ActiveHeap []project

func (h PausedHeap) Len() int           { return len(h) }
func (h ActiveHeap) Len() int           { return len(h) }
func (h PausedHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h ActiveHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h PausedHeap) Less(i, j int) bool { return h[i].capital < h[j].capital }
func (h ActiveHeap) Less(i, j int) bool { return h[i].profit > h[j].profit }
func (h *PausedHeap) Pop() any          { ret := (*h)[len(*h)-1]; *h = (*h)[0 : len(*h)-1]; return ret }
func (h *ActiveHeap) Pop() any          { ret := (*h)[len(*h)-1]; *h = (*h)[0 : len(*h)-1]; return ret }
func (h *PausedHeap) Push(i any)        { *h = append(*h, i.(project)) }
func (h *ActiveHeap) Push(i any)        { *h = append(*h, i.(project)) }

func findMaximizedCapital2(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	pausedProjects := make(PausedHeap, 0, n)
	activeProjects := make(ActiveHeap, 0, n)
	for i := range n {
		if capital[i] > w {
			pausedProjects = append(pausedProjects, project{capital[i], profits[i], i})
		} else {
			activeProjects = append(activeProjects, project{capital[i], profits[i], i})
		}
	}
	heap.Init(&pausedProjects)
	heap.Init(&activeProjects)
	for k > 0 && len(activeProjects) > 0 {
		topProject := heap.Pop(&activeProjects).(project)
		w += topProject.profit
		for pausedProjects.Len() > 0 && pausedProjects[0].capital <= w {
			heap.Push(&activeProjects, heap.Pop(&pausedProjects))
		}
		k--
	}
	return w
}
