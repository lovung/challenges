package jul2024

import (
	"math"

	"github.com/emirpasic/gods/v2/queues/linkedlistqueue"
)

// https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/
type edge struct{ t, w int }

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	adj := make([][]edge, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], edge{e[1], e[2]})
		adj[e[1]] = append(adj[e[1]], edge{e[0], e[2]})
	}
	minCityCnt := math.MaxInt
	res := -1
	for i := range n {
		cnt := bfsFindTheCity(adj, i, distanceThreshold)
		if cnt <= minCityCnt {
			minCityCnt = cnt
			res = i
		}
	}
	return res
}

func bfsFindTheCity(adj [][]edge, from int, threshold int) int {
	type item struct{ node, remainDistance int }
	visited := make(map[int]int)
	q := linkedlistqueue.New[item]()
	q.Enqueue(item{from, threshold})
	for !q.Empty() {
		curNode, _ := q.Dequeue()
		if val, ok := visited[curNode.node]; ok && val >= curNode.remainDistance {
			continue
		}
		visited[curNode.node] = curNode.remainDistance
		for _, adjEdge := range adj[curNode.node] {
			if adjEdge.w <= curNode.remainDistance {
				q.Enqueue(item{adjEdge.t, curNode.remainDistance - adjEdge.w})
			}
		}
	}
	return len(visited) - 1
}

func findTheCity_floyd(n int, edges [][]int, thresh int) int {
	// REPRESENT THE GRAPH AS AN ADJACENCY MATRIX
	g := getAdjMat(n, edges)

	// FOR EACH INTERMEDIATE NODE K
	for k := range n {
		// TRY EACH PAIR
		for i := range n {
			for j := range n {
				g[i][j] = min(g[i][j], g[i][k]+g[k][j])
			}
		}
	}

	// GET THE CITY WITH THE FEWEST NUMBER OF LINKS
	city, links := -1, int(math.Pow(2, 64))
	for i := range n {
		l := 0
		for j := range n {
			if g[i][j] <= thresh {
				l++
			}
		}
		if city == -1 || l <= links {
			city, links = i, l
		}
	}
	return city
}

func getAdjMat(n int, edges [][]int) [][]int {
	g := make([][]int, n)

	for i := range n {
		g[i] = make([]int, n)
	}

	// DEFAULT TO 2**64
	for i := range n {
		for j := range n {
			g[i][j] = int(math.Pow(2, 60))
		}
	}

	// EDGES
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g[u][v] = w
		g[v][u] = w
	}

	// SELF-LOOPS
	for i := range n {
		g[i][i] = 0
	}

	return g
}
