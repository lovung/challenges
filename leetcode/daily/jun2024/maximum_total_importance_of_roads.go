package jun2024

import "slices"

// https://leetcode.com/problems/maximum-total-importance-of-roads/description
func maximumImportance(n int, roads [][]int) int64 {
	adj := make([][]int, n)
	for _, r := range roads {
		adj[r[0]] = append(adj[r[0]], r[1])
		adj[r[1]] = append(adj[r[1]], r[0])
	}
	cities := make([]int, 0, n)
	for i := range n {
		cities = append(cities, i)
	}
	slices.SortFunc(cities, func(a, b int) int {
		return len(adj[b]) - len(adj[a])
	})
	res := int64(0)
	for i := range n {
		res += int64((n - i) * len(adj[cities[i]]))
	}
	return res
}
