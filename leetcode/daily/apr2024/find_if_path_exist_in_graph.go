package apr2024

import "github.com/lovung/ds/graphs"

// https://leetcode.com/problems/find-if-path-exists-in-graph/
func validPath(n int, edges [][]int, source int, destination int) bool {
	uf := graphs.NewUnionFind(n)
	for _, e := range edges {
		uf.Union(e[0], e[1])
	}
	return uf.Find(source) == uf.Find(destination)
}
