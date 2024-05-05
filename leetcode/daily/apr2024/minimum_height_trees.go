package apr2024

import (
	"github.com/lovung/ds/queue"
)

// https://leetcode.com/problems/minimum-height-trees
func findMinHeightTrees(n int, edges [][]int) []int {
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
	leaves := queue.NewSimpleQueue[int]()
	for i := 0; i < n; i++ {
		if len(neighbors[i]) == 1 {
			leaves.EnQueue(i)
		}
	}

	remainNode := n
	for remainNode > 2 {
		newLeaves := queue.NewSimpleQueue[int]()
		remainNode -= leaves.Len()
		for leaves.Len() > 0 {
			leaf, _ := leaves.DeQueue()
			for k := range neighbors[leaf] {
				delete(neighbors[k], leaf)
				if len(neighbors[k]) == 1 {
					newLeaves.EnQueue(k)
				}
			}
		}
		leaves = newLeaves
	}

	q := leaves.(*queue.SimpleQueue[int])
	return []int(*q)
}
