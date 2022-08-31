package connectedgraph

import (
	"container/heap"

	"github.com/lovung/ds/heaps"
	"github.com/lovung/ds/maths"
	"github.com/lovung/ds/queue"
)

// Link: https://leetcode.com/problems/minimum-height-trees/
func findMinHeightTrees2(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	neighbors := make([]map[int]struct{}, n)
	for i := range edges {
		if neighbors[edges[i][0]] == nil {
			neighbors[edges[i][0]] = make(map[int]struct{})
		}
		neighbors[edges[i][0]][edges[i][1]] = struct{}{}
		if neighbors[edges[i][1]] == nil {
			neighbors[edges[i][1]] = make(map[int]struct{})
		}
		neighbors[edges[i][1]][edges[i][0]] = struct{}{}
	}
	leaves := queue.NewQueue[int]()
	for i := 0; i < n; i++ {
		if len(neighbors[i]) == 1 {
			leaves.Push(i)
		}
	}

	remainNode := n
	for remainNode > 2 {
		newLeaves := queue.NewQueue[int]()
		remainNode -= leaves.Len()
		for leaves.Len() > 0 {
			left := leaves.Pop()
			for k := range neighbors[left] {
				delete(neighbors[k], left)
				if len(neighbors[k]) == 1 {
					newLeaves.Push(k)
				}
			}

		}
		leaves = newLeaves
	}
	return []int(leaves)
}

// Link: https://leetcode.com/problems/minimum-height-trees/
func findMinHeightTrees1(n int, edges [][]int) []int {
	links := make([][]int, n)
	for i := range edges {
		links[edges[i][0]] = append(links[edges[i][0]], edges[i][1])
		links[edges[i][1]] = append(links[edges[i][1]], edges[i][0])
	}
	minHeight := n
	minHeap := heaps.MinHeapWithValue[int]{}
	for i := 0; i < n; i++ {
		// dfs instead
		mark := make([]bool, n)
		height := findMinHeightTreesBFS(n, links, i, mark)
		if height < minHeight {
			minHeight = height
		}
		heap.Push(&minHeap, &heaps.HeapItem[int]{
			Ref:   height,
			Value: i,
		})
	}

	res := make([]int, 0)
	for item := heap.Pop(&minHeap).(*heaps.HeapItem[int]); item.Ref == minHeight; {
		res = append(res, item.Value.(int))
		if minHeap.Len() == 0 {
			break
		}
		item = heap.Pop(&minHeap).(*heaps.HeapItem[int])
	}
	return res
}

// return height
func findMinHeightTreesBFS(n int, links [][]int, start int, mark []bool) int {
	if mark[start] {
		return 0
	}
	if len(links[start]) == 0 {
		return 1
	}
	mark[start] = true
	subHeight := make([]int, len(links[start]))
	for i := range links[start] {
		subHeight = append(subHeight, findMinHeightTreesBFS(n, links, links[start][i], mark)+1)
	}
	return maths.Max(subHeight...)
}
