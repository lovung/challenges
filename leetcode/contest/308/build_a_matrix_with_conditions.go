package contest

import (
	llq "github.com/emirpasic/gods/v2/queues/linkedlistqueue"
)

// Link: https://leetcode.com/problems/build-a-matrix-with-conditions/
func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	rowOrder := toposortKahnAlgo(k, rowConditions)
	if len(rowOrder) != k {
		return [][]int{}
	}
	colOrder := toposortKahnAlgo(k, colConditions)
	if len(colOrder) != k {
		return [][]int{}
	}
	result := make([][]int, k)
	colIndex := make([]int, k+1)
	for i, c := range colOrder {
		colIndex[c] = i
	}
	for i, r := range rowOrder {
		result[i] = make([]int, k)
		result[i][colIndex[r]] = r
	}
	return result
}

// Ref: https://en.wikipedia.org/wiki/Topological_sorting
// and: https://leetcode.com/problems/build-a-matrix-with-conditions/discuss/2492865/Khan's-Algo-or-Topological-Sort-or-BFS-or-C%2B%2B-or-Short-and-Easy-to-Understand-Code
func toposortKahnAlgo(k int, con [][]int) []int {
	parentCnt := make([]int, k+1)
	// Adjacency list: https://en.wikipedia.org/wiki/Adjacency_list
	adjList := make([][]int, k+1)
	for i := range con {
		parentCnt[con[i][1]]++
		adjList[con[i][0]] = append(adjList[con[i][0]], con[i][1])
	}
	res := make([]int, 0, k)
	s := llq.New[int]()
	for i := 1; i <= k; i++ {
		if parentCnt[i] == 0 {
			s.Enqueue(i)
		}
	}
	// while S is not empty do
	for !s.Empty() {
		// remove a node n from S
		n, _ := s.Dequeue()
		// add n to L
		res = append(res, n)
		// for each node m with an edge e from n to m do
		for _, m := range adjList[n] {
			// remove edge e from the graph
			parentCnt[m]--
			// if m has no other incoming edges then
			if parentCnt[m] == 0 {
				// insert m into S
				s.Enqueue(m)
			}
		}
	}
	return res
}
