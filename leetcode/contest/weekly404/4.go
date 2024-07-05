package weekly404

import "github.com/emirpasic/gods/v2/queues/linkedlistqueue"

func minimumDiameterAfterMerge_WA(edges1 [][]int, edges2 [][]int) int {
	n := len(edges1) + 1
	m := len(edges2) + 1
	return 1 +
		findMinHeightTrees(n, edges1) +
		findMinHeightTrees(m, edges2)
}

func findMinHeightTrees(n int, edges [][]int) int {
	if n <= 2 {
		return n - 1
	}
	neighbors := buildAdj(n, edges)
	leaves := linkedlistqueue.New[int]()
	for i := 0; i < n; i++ {
		if len(neighbors[i]) == 1 {
			leaves.Enqueue(i)
		}
	}

	remainNode := n
	height := 0
	for remainNode > 2 {
		newLeaves := linkedlistqueue.New[int]()
		remainNode -= leaves.Size()
		for leaves.Size() > 0 {
			left, _ := leaves.Dequeue()
			for k := range neighbors[left] {
				delete(neighbors[k], left)
				if len(neighbors[k]) == 1 {
					newLeaves.Enqueue(k)
				}
			}
		}
		leaves = newLeaves
		height++
	}
	if remainNode == 2 {
		height++
	}
	return height
}

func buildAdj(n int, edges [][]int) []map[int]struct{} {
	neighbors := make([]map[int]struct{}, n)
	for _, e := range edges {
		if neighbors[e[0]] == nil {
			neighbors[e[0]] = make(map[int]struct{})
		}
		neighbors[e[0]][e[1]] = struct{}{}
		if neighbors[e[1]] == nil {
			neighbors[e[1]] = make(map[int]struct{})
		}
		neighbors[e[1]][e[0]] = struct{}{}
	}
	return neighbors
}
