package contest

import "github.com/lovung/ds/queue"

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
	for j := range colOrder {
		colIndex[colOrder[j]] = j
	}
	for i := 0; i < k; i++ {
		if result[i] == nil {
			result[i] = make([]int, k)
		}
		result[i][colIndex[rowOrder[i]]] = rowOrder[i]
	}
	return result
}

// Ref: https://en.wikipedia.org/wiki/Topological_sorting
// and: https://leetcode.com/problems/build-a-matrix-with-conditions/discuss/2492865/Khan's-Algo-or-Topological-Sort-or-BFS-or-C%2B%2B-or-Short-and-Easy-to-Understand-Code
func toposortKahnAlgo(k int, con [][]int) []int {
	parentCnt := make([]int, k+1)
	childrenList := make([][]int, k+1)
	for i := range con {
		parentCnt[con[i][1]]++
		childrenList[con[i][0]] = append(childrenList[con[i][0]], con[i][1])
	}
	res := make([]int, 0, k)
	s := queue.NewQueue[int]()
	for i := 1; i <= k; i++ {
		if parentCnt[i] == 0 {
			s.Push(i)
		}
	}
	// while S is not empty do
	for s.Len() > 0 {
		// remove a node n from S
		n := s.Pop()
		// add n to L
		res = append(res, n)
		// for each node m with an edge e from n to m do
		for _, m := range childrenList[n] {
			// remove edge e from the graph
			parentCnt[m]--
			// f m has no other incoming edges then
			if parentCnt[m] == 0 {
				// insert m into S
				s.Push(m)
			}
		}
	}
	return res
}
