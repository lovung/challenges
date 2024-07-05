package jun2024

import (
	"slices"
	"sort"
)

// https://leetcode.com/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/
func getAncestors(n int, edges [][]int) [][]int {
	parents := make([][]int, n)
	for _, e := range edges {
		parents[e[1]] = append(parents[e[1]], e[0])
	}
	for i := range n {
		if len(parents[i]) == 0 {
			parents[i] = []int{}
		} else {
			sort.Ints(parents[i])
		}
	}
	visited := make([]bool, n)
	for i := range n {
		visitNode(i, parents, visited)
	}

	return parents
}

func visitNode(i int, parents [][]int, visited []bool) []int {
	if visited[i] {
		return parents[i]
	}
	visited[i] = true
	ancesters := make([]int, 0)
	for _, p := range parents[i] {
		mergeAncestor(visitNode(p, parents, visited), &ancesters)
	}
	mergeAncestor(ancesters, &parents[i])
	return parents[i]
}

func mergeAncestor(ancestors []int, parents *[]int) {
	for _, a := range ancestors {
		if idx, ok := slices.BinarySearch(*parents, a); !ok {
			*parents = slices.Insert(*parents, idx, a)
		}
	}
}
