package oct2022

import "github.com/lovung/ds/queue"

// Link: https://leetcode.com/problems/bus-routes/

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

	type QueueItem struct {
		station int
		level   int
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

func numBusesToDestination2(routes [][]int, source int, target int) int {
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

	// routeIndex -> adjecentRouteIndex
	// If want to remove duplicate,
	// adjacentRouteMap := make(map[int]map[int]struct{})
	// But we have visited already to handle it
	adjacentRouteMap := make(map[int][]int)
	for _, v := range busIndexMap {
		if len(v) <= 1 {
			// skip the alone station
			continue
		}
		for i := range v {
			// for j := range v {
			// 	if i == j {
			// 		continue
			// 	}
			// 	if adjacentRouteMap[v[i]] == nil {
			// 		adjacentRouteMap[v[i]] = make(map[int]struct{})
			// 	}
			// 	adjacentRouteMap[v[i]][v[j]] = struct{}
			// }
			adjacentRouteMap[v[i]] = append(adjacentRouteMap[v[i]], v[:i]...)
			adjacentRouteMap[v[i]] = append(adjacentRouteMap[v[i]], v[i+1:]...)
		}
	}

	type QueueItem struct {
		route int
		level int
	}
	q := queue.NewQueue[*QueueItem]()
	visited := make(map[int]bool)
	for _, routeFromSource := range busIndexMap[source] {
		q.Push(&QueueItem{routeFromSource, 1})
		visited[routeFromSource] = true
	}
	routesCanReachTarget := make(map[int]bool)
	for _, routeCanReachTarget := range busIndexMap[target] {
		routesCanReachTarget[routeCanReachTarget] = true
	}

	for q.Len() > 0 {
		item := q.Pop()
		if routesCanReachTarget[item.route] {
			return item.level
		}
		for _, routeIndex := range adjacentRouteMap[item.route] {
			if visited[routeIndex] {
				continue
			}
			visited[routeIndex] = true
			q.Push(&QueueItem{routeIndex, item.level + 1})
		}
	}

	return -1
}
