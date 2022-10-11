package oct2022

import "github.com/lovung/ds/queue"

// Link: https://leetcode.com/problems/bus-routes/

// If only need to return true / false, we can use UnionFind to group all
// bus stops in the same route together. Check if the `source` and `target` is
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
		stop  int
		level int
	}
	q := queue.NewQueue[*QueueItem]()
	visited := make(map[int]bool)
	q.Push(&QueueItem{source, 0})
	visited[source] = true

	for q.Len() > 0 {
		item := q.Pop()
		for _, routeIndex := range busIndexMap[item.stop] {
			for _, stopNum := range routes[routeIndex] {
				if visited[stopNum] {
					continue
				}
				if stopNum == target {
					return item.level + 1
				}
				visited[stopNum] = true
				q.Push(&QueueItem{stopNum, item.level + 1})
			}
		}
	}

	return -1
}

// R is number of routes
// S is max number of stops in 1 route
func numBusesToDestination2(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	// Build adjacent map:
	// Time: O(R*S), Space: O(R*S)
	busIndexMap := buildBusIndexMap(routes)

	// Check if no routes go from the source or pass the target
	if len(busIndexMap[source]) == 0 || len(busIndexMap[target]) == 0 {
		return -1
	}

	// Time: O(R*S), Space: O(R)
	adjacentRouteMap := buildAdjacentRouteMap(busIndexMap)

	// Time: O(R), Space: O(R)
	routesCanReachTarget := buildRoutesCanReachTarget(busIndexMap, target)

	// Time: O(R), Space: O(R)
	return bfsNumBusesToDestination(
		busIndexMap,
		adjacentRouteMap,
		source,
		routesCanReachTarget,
	)
}

// Time: O(R), Space: O(R)
func bfsNumBusesToDestination(
	busIndexMap map[int][]int,
	adjacentRouteMap map[int][]int,
	source int,
	routesCanReachTarget map[int]bool,
) int {
	type QueueItem struct {
		route int
		level int
	}
	q := queue.NewQueue[*QueueItem]()
	visited := make(map[int]bool)

	// Small time
	for _, routeFromSource := range busIndexMap[source] {
		q.Push(&QueueItem{routeFromSource, 1})
		visited[routeFromSource] = true
	}

	// O(R)
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

// Time: O(R*S), Space: O(R*S)
func buildBusIndexMap(routes [][]int) map[int][]int {
	busIndexMap := make(map[int][]int)
	for i := range routes {
		for j := range routes[i] {
			busIndexMap[routes[i][j]] = append(busIndexMap[routes[i][j]], i)
		}
	}
	return busIndexMap
}

// Time: O(R), Space: O(R)
func buildRoutesCanReachTarget(busIndexMap map[int][]int, target int) map[int]bool {
	routesCanReachTarget := make(map[int]bool)
	for _, routeCanReachTarget := range busIndexMap[target] {
		routesCanReachTarget[routeCanReachTarget] = true
	}
	return routesCanReachTarget
}

// If want to remove duplicate,
// adjacentRouteMap := make(map[int]map[int]struct{})
// But we have visited already to handle it
// Map from routeIndex to adjecentRouteIndex
// Time: O(R*S), Space: O(R)
func buildAdjacentRouteMap(busIndexMap map[int][]int) map[int][]int {
	adjacentRouteMap := make(map[int][]int)
	for _, v := range busIndexMap {
		if len(v) <= 1 {
			// skip the alone stop
			continue
		}
		for i := range v {
			adjacentRouteMap[v[i]] = append(adjacentRouteMap[v[i]], v[:i]...)
			adjacentRouteMap[v[i]] = append(adjacentRouteMap[v[i]], v[i+1:]...)
		}
	}
	return adjacentRouteMap
}
