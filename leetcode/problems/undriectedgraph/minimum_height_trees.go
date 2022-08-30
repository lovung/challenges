package undriectedgraph

import "github.com/lovung/ds/queue"

// Link: https://leetcode.com/problems/minimum-height-trees/
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
