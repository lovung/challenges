package apr2024

import . "github.com/lovung/ds/queue"

// https://leetcode.com/problems/sum-of-distances-in-tree/
func sumOfDistancesInTree1(n int, edges [][]int) []int {
	if n <= 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{1, 1}
	}
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	// don't have leaves
	leaveCnt := make([]int, n)
	simAdj := make([][]int, n)
	root := make([]int, 0, n)
	for i := range adj {
		if len(adj[i]) > 1 {
			root = append(root, i)
		}
		for _, k := range adj[i] {
			if len(adj[k]) > 1 {
				simAdj[i] = append(simAdj[i], k)
			} else {
				leaveCnt[i]++
			}
		}
	}
	res := make([]int, n)
	for i := range root {
		res[root[i]] = bfsFrom(n, root[i], simAdj, leaveCnt)
	}
	for i := range n {
		if res[i] == 0 {
			res[i] = res[adj[i][0]] + (n - 2)
		}
	}
	return res
}

func bfsFrom(n, root int, adj [][]int, leaveCnt []int) int {
	visited := make([]bool, n)
	type step struct {
		val, cost int
	}
	res := 0
	q := NewCircularQueue[step](n)
	q.EnQueue(step{root, 0})
	for q.Len() > 0 {
		cur, _ := q.DeQueue()
		res += cur.cost + leaveCnt[cur.val]*(cur.cost+1)
		visited[cur.val] = true
		for _, near := range adj[cur.val] {
			if !visited[near] {
				q.EnQueue(step{near, cur.cost + 1})
			}
		}
	}
	return res
}
