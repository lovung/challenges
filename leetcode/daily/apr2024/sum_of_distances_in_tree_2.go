package apr2024

// import (
// 	. "github.com/lovung/ds/queue"
// 	. "github.com/lovung/ds/stack"
// )

// // https://leetcode.com/problems/sum-of-distances-in-tree/
// func sumOfDistancesInTree(n int, edges [][]int) []int {
// 	if n <= 1 {
// 		return []int{0}
// 	}
// 	if n == 2 {
// 		return []int{1, 1}
// 	}
// 	adjMap := make([]map[int]bool, n)
// 	adj := make([][]int, n)
// 	for i := range adjMap {
// 		adjMap[i] = make(map[int]bool)
// 	}
// 	for _, e := range edges {
// 		adj[e[0]] = append(adj[e[0]], e[1])
// 		adj[e[1]] = append(adj[e[1]], e[0])
// 		adjMap[e[0]][e[1]] = true
// 		adjMap[e[1]][e[0]] = true
// 	}
// 	// don't have leaves
// 	leaveCnt := make([]int, n)
// 	leaveCost := make([]int, n)
// 	q := NewCircularQueue[int](n)
// 	// type calcLater struct {
// 	// 	val                   int
// 	// 	targetCnt, targetCost int
// 	// }
// 	s := NewArrayStack[int](n)
// 	for i := range adjMap {
// 		if len(adjMap[i]) == 1 {
// 			q.EnQueue(i)
// 		}
// 	}
// 	res := make([]int, n)
// 	merge := make([]int, n)
// 	for i := range merge {
// 		merge[i] = -1
// 	}
// 	for remain := n; remain > 2; remain-- {
// 		leaf, ok := q.DeQueue()
// 		if !ok {
// 			break
// 		}
// 		for k := range adjMap[leaf] {
// 			// leaveCntK, leaveCostK := leaveCnt[k], leaveCost[k]
// 			leaveCnt[k] += leaveCnt[leaf] + 1
// 			// leaveCost[k] += 1 + leaveCost[leaf] + leaveCnt[leaf]
// 			leaveCost[k] += calcCost(leaveCost[leaf], leaveCnt[leaf])
// 			// leaveCnt[leaf] += leaveCntK + 1
// 			// leaveCost[leaf] += 1 + leaveCostK + leaveCntK
// 			delete(adjMap[k], leaf)
// 			merge[leaf] = k
// 			if len(adjMap[k]) == 1 {
// 				q.EnQueue(k)
// 			}
// 			s.Push(leaf)
// 		}
// 	}
// 	for i := range merge {
// 		if merge[i] == -1 {
// 			for k := range adjMap[i] {
// 				res[i] = leaveCost[i] + calcCost(leaveCost[k], leaveCnt[k])
// 			}
// 		}
// 	}
// 	for s.Len() > 0 {
// 		i := s.Pop()
// 		k := merge[i]
// 		reverseCnt := n - 1 - leaveCnt[k]
// 		reverseCost := res[k] - leaveCost[k]
// 		res[i] = leaveCost[i] + calcCost(reverseCost, reverseCnt)
// 	}
// 	return res
// }

// func calcCost(cost, cnt int) int {
// 	return 1 + cost + cnt
// }
