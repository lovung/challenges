package oct2022

import "github.com/lovung/ds/queue"

// Link: https://leetcode.com/problems/bus-routes/
type QueueItem struct {
	station int
	level   int
}

// If only need to return true / false, we can use UnionFind to group all
// bus stations in the same route together. Check if the `source` and `target` is
// in the same group after that.

// Let's go with BFS
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	// Build adjacent map: O(m*n)
	busIndexMap := make(map[int][]int)
	for i := range routes {
		for j := range routes[i] {
			busIndexMap[routes[i][j]] = append(busIndexMap[routes[i][j]], i)
		}
	}

	if len(busIndexMap[source]) == 0 || len(busIndexMap[target]) == 0 {
		return -1
	}

	q := queue.NewQueue[*QueueItem]()
	visited := make(map[int]bool)
	q.Push(&QueueItem{source, 0})
	visited[source] = true

	for q.Len() > 0 {
		item := q.Pop()
		for _, routeIndex := range busIndexMap[item.station] {
			for _, stationNum := range routes[routeIndex] {
				if visited[stationNum] {
					continue
				}
				if stationNum == target {
					return item.level + 1
				}
				visited[stationNum] = true
				q.Push(&QueueItem{stationNum, item.level + 1})
			}
		}
	}

	return -1
}
